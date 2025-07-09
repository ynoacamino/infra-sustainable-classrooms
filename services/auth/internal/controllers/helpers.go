package controllers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/netip"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	auth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"golang.org/x/crypto/bcrypt"
)

// getSessionExpirationPG returns session expiration time in PostgreSQL format
func (s *authsrvc) getSessionExpirationPG() pgtype.Timestamptz {
	expiresAt := time.Now().Add(s.cfg.SessionDuration)
	return pgtype.Timestamptz{
		Time:  expiresAt,
		Valid: true,
	}
}

// generateSessionToken generates a secure session token
func (s *authsrvc) generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// generateBackupCodes generates a specified number of backup codes
func (s *authsrvc) generateBackupCodes(count int) ([]string, error) {
	codes := make([]string, count)
	for i := range count {
		code, err := s.generateRandomCode(8)
		if err != nil {
			return nil, err
		}
		codes[i] = code
	}
	return codes, nil
}

// generateRandomCode generates a random alphanumeric code of specified length
func (s *authsrvc) generateRandomCode(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes), nil
}

// hashBackupCode hashes a backup code using bcrypt
func (s *authsrvc) hashBackupCode(code string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Sprintf("fallback_%s", code)
	}
	return string(hash)
}

// verifyBackupCode verifies a backup code against its hash
func (s *authsrvc) verifyBackupCode(code, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(code))
	return err == nil
}

// getDeviceInfoText extracts text field from device info
func (s *authsrvc) getDeviceInfoText(deviceInfo *auth.DeviceInfo, field string) pgtype.Text {
	var value string
	if deviceInfo != nil {
		switch field {
		case "user_agent":
			if deviceInfo.UserAgent != nil {
				value = *deviceInfo.UserAgent
			}
		case "device_id":
			if deviceInfo.DeviceID != nil {
				value = *deviceInfo.DeviceID
			}
		case "platform":
			if deviceInfo.Platform != nil {
				value = *deviceInfo.Platform
			}
		}
	}

	return pgtype.Text{
		String: value,
		Valid:  value != "",
	}
}

// getDeviceInfoIP extracts IP address from device info
func (s *authsrvc) getDeviceInfoIP(deviceInfo *auth.DeviceInfo) *netip.Addr {
	if deviceInfo == nil || deviceInfo.IPAddress == nil {
		return nil
	}

	addr, err := netip.ParseAddr(*deviceInfo.IPAddress)
	if err != nil {
		return nil
	}

	return &addr
}

func (s *authsrvc) validateSession(ctx context.Context, sessionToken string) (*authdb.GetSessionByTokenRow, error) {
	if sessionToken == "" {
		return nil, fmt.Errorf("session token is required")
	}

	sessionData, err := s.sessionRepo.GetSessionByToken(ctx, sessionToken)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if !sessionData.IsActive {
		return nil, fmt.Errorf("session is inactive")
	}

	if sessionData.ExpiresAt.Valid && sessionData.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("session already expired")
	}

	return &sessionData, nil
}

func (s *authsrvc) truncateToken(token string) string {
	if len(token) <= 10 {
		return token[:len(token)/2] + "..."
	}
	return token[:10] + "..."
}
