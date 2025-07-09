package controllers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pquerna/otp/totp"
	auth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/mappers"
	"goa.design/clue/log"
)

func (s *authsrvc) RefreshSession(ctx context.Context, p *auth.RefreshSessionPayload) (res *auth.SimpleResponse, err error) {
	log.Printf(ctx, "auth.RefreshSession for session token")

	sessionData, err := s.validateSession(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Session validation failed: %v", err)
		return nil, auth.Unauthorized("invalid or expired session")
	}

	newExpiration := s.getSessionExpirationPG()

	refreshedSession, err := s.sessionRepo.RefreshSession(ctx, authdb.RefreshSessionParams{
		SessionToken: p.SessionToken,
		ExpiresAt:    newExpiration,
	})
	if err != nil {
		log.Printf(ctx, "Error refreshing session: %v", err)
		return nil, auth.ServiceUnavailable("failed to refresh session")
	}

	res = &auth.SimpleResponse{
		Success:      true,
		Message:      "Session refreshed successfully",
		SessionToken: &refreshedSession.SessionToken,
	}

	log.Printf(ctx, "Successfully refreshed session for user ID: %d, new expiry: %v",
		sessionData.UserID, newExpiration.Time)
	return res, nil
}

func (s *authsrvc) ValidateUser(ctx context.Context, p *auth.ValidateUserPayload) (res *auth.UserValidationResponse, err error) {
	log.Printf(ctx, "auth.ValidateUser for session: %s", s.truncateToken(p.SessionToken))

	sessionData, err := s.validateSession(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Session validation failed: %v", err)
		return &auth.UserValidationResponse{
			Valid: false,
		}, nil
	}

	user, err := s.userRepo.GetUserByID(ctx, sessionData.UserID)
	if err != nil {
		log.Printf(ctx, "User not found for session: %v", err)
		return &auth.UserValidationResponse{
			Valid: false,
		}, nil
	}

	if !user.IsVerified {
		log.Printf(ctx, "User not verified: %d", user.ID)
	}

	err = s.sessionRepo.UpdateSessionAccess(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Warning: failed to update session access: %v", err)
	}

	apiUser := mappers.UserDBToAPI(&user)

	sessionForAPI := authdb.Session{
		ID:           sessionData.ID,
		UserID:       sessionData.UserID,
		SessionToken: sessionData.SessionToken,
		ExpiresAt:    sessionData.ExpiresAt,
		CreatedAt:    sessionData.CreatedAt,
		LastAccessed: sessionData.LastAccessed,
		IsActive:     sessionData.IsActive,
		UserAgent:    sessionData.UserAgent,
		IpAddress:    sessionData.IpAddress,
		DeviceID:     sessionData.DeviceID,
		Platform:     sessionData.Platform,
	}
	apiSession := mappers.SessionDBToAPI(&sessionForAPI)

	res = &auth.UserValidationResponse{
		Valid:   true,
		User:    apiUser,
		Session: apiSession,
	}

	log.Printf(ctx, "Successfully validated user: %d, session: %s",
		user.ID, s.truncateToken(p.SessionToken))
	return res, nil
}

func (s *authsrvc) GetUserByID(ctx context.Context, p *auth.GetUserByIDPayload) (res *auth.User, err error) {
	log.Printf(ctx, "auth.GetUserByID for user ID: %d", p.UserID)

	user, err := s.userRepo.GetUserByID(ctx, p.UserID)
	if err != nil {
		log.Printf(ctx, "User not found with ID: %d, error: %v", p.UserID, err)
		return nil, auth.UserNotFound("user not found")
	}

	apiUser := mappers.UserDBToAPI(&user)

	log.Printf(ctx, "Successfully retrieved user: %d (%s)", user.ID, user.Identifier)
	return apiUser, nil
}

func (s *authsrvc) Logout(ctx context.Context, p *auth.LogoutPayload) (res *auth.SimpleResponse, err error) {
	log.Printf(ctx, "auth.Logout for session token")

	sessionData, err := s.validateSession(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Session validation failed during logout: %v", err)
		return &auth.SimpleResponse{
			Success: true,
			Message: "Logout successful",
		}, nil
	}

	err = s.sessionRepo.DeactivateSession(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Error deactivating session: %v", err)
		return nil, auth.ServiceUnavailable("failed to deactivate session")
	}

	res = &auth.SimpleResponse{
		Success: true,
		Message: "Logout successful",
	}

	log.Printf(ctx, "Successfully logged out user ID: %d, session: %s",
		sessionData.UserID, s.truncateToken(p.SessionToken))
	return res, nil
}

func (s *authsrvc) GetUserProfile(ctx context.Context, p *auth.GetUserProfilePayload) (res *auth.User, err error) {
	log.Printf(ctx, "auth.GetUserProfile for session token")

	sessionData, err := s.validateSession(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Session validation failed: %v", err)
		return nil, auth.Unauthorized("invalid or expired session")
	}

	user, err := s.userRepo.GetUserByID(ctx, sessionData.UserID)
	if err != nil {
		log.Printf(ctx, "User not found: %v", err)
		return nil, auth.UserNotFound("user not found")
	}

	err = s.sessionRepo.UpdateSessionAccess(ctx, p.SessionToken)
	if err != nil {
		log.Printf(ctx, "Warning: failed to update session access: %v", err)
	}

	apiUser := mappers.UserDBToAPI(&user)

	log.Printf(ctx, "Successfully retrieved profile for user: %d (%s)",
		user.ID, user.Identifier)
	return apiUser, nil
}

func (s *authsrvc) VerifyBackupCode(ctx context.Context, p *auth.VerifyBackupCodePayload) (res *auth.BackupCodeResponse, err error) {
	log.Printf(ctx, "auth.VerifyBackupCode for identifier: %s", p.Identifier)

	user, err := s.userRepo.GetUserByIdentifier(ctx, p.Identifier)
	if err != nil {
		log.Printf(ctx, "User not found: %v", err)
		return nil, auth.UserNotFound("user not found")
	}

	var userAPI *auth.User
	var sessionToken string
	var remainingCount int64

	// Usar transacción para verificar y marcar el código como usado atómicamente
	err = s.txManager.WithTx(ctx, func(tx pgx.Tx) error {
		txRepos := s.repoManager.WithTransaction(tx)

		backupCodes, err := txRepos.BackupCodeRepo.GetUserBackupCodes(ctx, user.ID)
		if err != nil {
			return fmt.Errorf("failed to retrieve backup codes: %w", err)
		}

		var validCodeHash string
		for _, code := range backupCodes {
			if s.verifyBackupCode(p.BackupCode, code.CodeHash) {
				validCodeHash = code.CodeHash
				break
			}
		}

		if validCodeHash == "" {
			return auth.InvalidOtp("invalid or already used backup code")
		}

		_, err = txRepos.BackupCodeRepo.UseBackupCode(ctx, authdb.UseBackupCodeParams{
			UserID:   user.ID,
			CodeHash: validCodeHash,
		})
		if err != nil {
			return fmt.Errorf("failed to update backup code: %w", err)
		}

		remainingCount, err = txRepos.BackupCodeRepo.CountAvailableBackupCodes(ctx, user.ID)
		if err != nil {
			return fmt.Errorf("failed to count remaining codes: %w", err)
		}

		userAPI = mappers.UserDBToAPI(&user)

		sessionToken, err = s.generateSessionToken()
		if err != nil {
			return fmt.Errorf("failed to generate session token: %w", err)
		}

		sessionParams := authdb.CreateSessionParams{
			UserID:       user.ID,
			SessionToken: sessionToken,
			ExpiresAt:    s.getSessionExpirationPG(),
			UserAgent:    s.getDeviceInfoText(p.DeviceInfo, "user_agent"),
			IpAddress:    s.getDeviceInfoIP(p.DeviceInfo),
			DeviceID:     s.getDeviceInfoText(p.DeviceInfo, "device_id"),
			Platform:     s.getDeviceInfoText(p.DeviceInfo, "platform"),
		}

		_, err = txRepos.SessionRepo.CreateSession(ctx, sessionParams)
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		return nil
	})

	if err != nil {
		log.Printf(ctx, "Error in backup code verification transaction: %v", err)
		return nil, auth.ServiceUnavailable("failed to verify backup code")
	}

	res = &auth.BackupCodeResponse{
		Success:        true,
		Message:        "Backup code verification successful",
		User:           userAPI,
		RemainingCodes: int(remainingCount),
		SessionToken:   sessionToken,
	}

	log.Printf(ctx, "Successfully verified backup code for user: %s, remaining codes: %d",
		p.Identifier, remainingCount)
	return res, nil
}

func (s *authsrvc) GenerateSecret(ctx context.Context, p *auth.GenerateSecretPayload) (res *auth.TOTPSecret, err error) {
	log.Printf(ctx, "auth.GenerateSecret for identifier: %s", p.Identifier)

	existingUser, err := s.userRepo.GetUserByIdentifier(ctx, p.Identifier)
	if err == nil && existingUser.ID != 0 {
		return nil, auth.InvalidInput(fmt.Sprintf("user with identifier %s already exists", p.Identifier))
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.cfg.TOTPIssuer,
		AccountName: p.Identifier,
		SecretSize:  32,
	})
	if err != nil {
		log.Printf(ctx, "Error generating TOTP secret: %v", err)
		return nil, auth.ServiceUnavailable("failed to generate TOTP secret")
	}

	backupCodes, err := s.generateBackupCodes(s.cfg.BackupCodesCount)
	if err != nil {
		log.Printf(ctx, "Error generating backup codes: %v", err)
		return nil, auth.ServiceUnavailable("failed to generate backup codes")
	}

	var user authdb.User

	// Usar transacción para crear usuario y backup codes atómicamente
	err = s.txManager.WithTx(ctx, func(tx pgx.Tx) error {
		// Crear repositorios transaccionales
		txRepos := s.repoManager.WithTransaction(tx)

		userParams := authdb.CreateUserParams{
			Identifier: p.Identifier,
			TotpSecret: key.Secret(),
			IsVerified: false,
			Metadata:   []byte("{}"),
		}

		user, err = txRepos.UserRepo.CreateUser(ctx, userParams)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		backupCodesParams := make([]authdb.CreateBackupCodesParams, len(backupCodes))
		for i, code := range backupCodes {
			backupCodesParams[i] = authdb.CreateBackupCodesParams{
				UserID:   user.ID,
				CodeHash: s.hashBackupCode(code),
			}
		}

		_, err = txRepos.BackupCodeRepo.CreateBackupCodes(ctx, backupCodesParams)
		if err != nil {
			return fmt.Errorf("failed to create backup codes: %w", err)
		}

		return nil
	})

	if err != nil {
		log.Printf(ctx, "Error in transaction: %v", err)
		return nil, auth.ServiceUnavailable("failed to create user with backup codes")
	}

	res = &auth.TOTPSecret{
		TotpURL:     key.URL(),
		BackupCodes: backupCodes,
		Issuer:      "Auth Service",
	}

	log.Printf(ctx, "Successfully generated TOTP secret for user: %s", p.Identifier)
	return res, nil
}

func (s *authsrvc) VerifyTOTP(ctx context.Context, p *auth.VerifyTOTPPayload) (res *auth.AuthResponse, err error) {
	log.Printf(ctx, "auth.VerifyTOTP for identifier: %s", p.Identifier)

	user, err := s.userRepo.GetUserByIdentifier(ctx, p.Identifier)
	if err != nil {
		log.Printf(ctx, "User not found: %v", err)
		return nil, auth.UserNotFound("user not found")
	}

	valid := totp.Validate(p.TotpCode, user.TotpSecret)
	if !valid {
		log.Printf(ctx, "Invalid TOTP code for user: %s", p.Identifier)
		return nil, auth.InvalidOtp("invalid TOTP code")
	}

	sessionToken, err := s.generateSessionToken()
	if err != nil {
		log.Printf(ctx, "Error generating session token: %v", err)
		return nil, auth.ServiceUnavailable("failed to generate session token")
	}

	sessionParams := authdb.CreateSessionParams{
		UserID:       user.ID,
		SessionToken: sessionToken,
		ExpiresAt:    s.getSessionExpirationPG(),
		UserAgent:    s.getDeviceInfoText(p.DeviceInfo, "user_agent"),
		IpAddress:    s.getDeviceInfoIP(p.DeviceInfo),
		DeviceID:     s.getDeviceInfoText(p.DeviceInfo, "device_id"),
		Platform:     s.getDeviceInfoText(p.DeviceInfo, "platform"),
	}

	session, err := s.sessionRepo.CreateSession(ctx, sessionParams)
	if err != nil {
		log.Printf(ctx, "Error creating session: %v", err)
		return nil, auth.ServiceUnavailable("failed to create session")
	}

	apiUser := mappers.UserDBToAPI(&user)

	res = &auth.AuthResponse{
		Success:      true,
		Message:      "Authentication successful",
		User:         apiUser,
		SessionToken: sessionToken,
	}

	log.Printf(ctx, "Successfully authenticated user: %s, session: %s",
		p.Identifier, s.truncateToken(session.SessionToken))
	return res, nil
}
