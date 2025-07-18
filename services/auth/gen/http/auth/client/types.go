// Code generated by goa v3.21.1, DO NOT EDIT.
//
// auth HTTP client types
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design/api
// -o ./services/auth/

package client

import (
	auth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	goa "goa.design/goa/v3/pkg"
)

// GenerateSecretRequestBody is the type of the "auth" service "GenerateSecret"
// endpoint HTTP request body.
type GenerateSecretRequestBody struct {
	// User identifier (username/email)
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
}

// VerifyTOTPRequestBody is the type of the "auth" service "VerifyTOTP"
// endpoint HTTP request body.
type VerifyTOTPRequestBody struct {
	// User identifier
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
	// 6-digit TOTP code from authenticator app
	TotpCode string `form:"totp_code" json:"totp_code" xml:"totp_code"`
	// Device information for security
	DeviceInfo *DeviceInfoRequestBody `form:"device_info,omitempty" json:"device_info,omitempty" xml:"device_info,omitempty"`
}

// VerifyBackupCodeRequestBody is the type of the "auth" service
// "VerifyBackupCode" endpoint HTTP request body.
type VerifyBackupCodeRequestBody struct {
	// User identifier
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
	// 8-character backup code
	BackupCode string `form:"backup_code" json:"backup_code" xml:"backup_code"`
	// Device information for security
	DeviceInfo *DeviceInfoRequestBody `form:"device_info,omitempty" json:"device_info,omitempty" xml:"device_info,omitempty"`
}

// GenerateSecretResponseBody is the type of the "auth" service
// "GenerateSecret" endpoint HTTP response body.
type GenerateSecretResponseBody struct {
	// TOTP URL in otpauth:// format for authenticator apps
	TotpURL *string `form:"totp_url,omitempty" json:"totp_url,omitempty" xml:"totp_url,omitempty"`
	// Backup recovery codes
	BackupCodes []string `form:"backup_codes,omitempty" json:"backup_codes,omitempty" xml:"backup_codes,omitempty"`
	// Service name for authenticator app
	Issuer *string `form:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
}

// VerifyTOTPResponseBody is the type of the "auth" service "VerifyTOTP"
// endpoint HTTP response body.
type VerifyTOTPResponseBody struct {
	// Authentication success status
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// Response message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// User information
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// VerifyBackupCodeResponseBody is the type of the "auth" service
// "VerifyBackupCode" endpoint HTTP response body.
type VerifyBackupCodeResponseBody struct {
	// Authentication success status
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// Response message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// User information
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	// Number of remaining backup codes
	RemainingCodes *int `form:"remaining_codes,omitempty" json:"remaining_codes,omitempty" xml:"remaining_codes,omitempty"`
}

// RefreshSessionResponseBody is the type of the "auth" service
// "RefreshSession" endpoint HTTP response body.
type RefreshSessionResponseBody struct {
	// Operation success status
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// Response message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LogoutResponseBody is the type of the "auth" service "Logout" endpoint HTTP
// response body.
type LogoutResponseBody struct {
	// Operation success status
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// Response message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Session token for cookie (when applicable)
	SessionToken *string `form:"session_token,omitempty" json:"session_token,omitempty" xml:"session_token,omitempty"`
}

// GetUserProfileResponseBody is the type of the "auth" service
// "GetUserProfile" endpoint HTTP response body.
type GetUserProfileResponseBody struct {
	// User unique identifier
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Phone number or email
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Account creation timestamp in milliseconds
	CreatedAt *int64 `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Last login timestamp in milliseconds
	LastLogin *int64 `form:"last_login,omitempty" json:"last_login,omitempty" xml:"last_login,omitempty"`
	// Account verification status
	IsVerified *bool `form:"is_verified,omitempty" json:"is_verified,omitempty" xml:"is_verified,omitempty"`
	// Additional user metadata
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// DeviceInfoRequestBody is used to define fields on request body types.
type DeviceInfoRequestBody struct {
	// Browser/app user agent
	UserAgent *string `form:"user_agent,omitempty" json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// Client IP address
	IPAddress *string `form:"ip_address,omitempty" json:"ip_address,omitempty" xml:"ip_address,omitempty"`
	// Unique device identifier
	DeviceID *string `form:"device_id,omitempty" json:"device_id,omitempty" xml:"device_id,omitempty"`
	// Platform (web, ios, android)
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	// User unique identifier
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Phone number or email
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Account creation timestamp in milliseconds
	CreatedAt *int64 `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Last login timestamp in milliseconds
	LastLogin *int64 `form:"last_login,omitempty" json:"last_login,omitempty" xml:"last_login,omitempty"`
	// Account verification status
	IsVerified *bool `form:"is_verified,omitempty" json:"is_verified,omitempty" xml:"is_verified,omitempty"`
	// Additional user metadata
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// NewGenerateSecretRequestBody builds the HTTP request body from the payload
// of the "GenerateSecret" endpoint of the "auth" service.
func NewGenerateSecretRequestBody(p *auth.GenerateSecretPayload) *GenerateSecretRequestBody {
	body := &GenerateSecretRequestBody{
		Identifier: p.Identifier,
	}
	return body
}

// NewVerifyTOTPRequestBody builds the HTTP request body from the payload of
// the "VerifyTOTP" endpoint of the "auth" service.
func NewVerifyTOTPRequestBody(p *auth.VerifyTOTPPayload) *VerifyTOTPRequestBody {
	body := &VerifyTOTPRequestBody{
		Identifier: p.Identifier,
		TotpCode:   p.TotpCode,
	}
	if p.DeviceInfo != nil {
		body.DeviceInfo = marshalAuthDeviceInfoToDeviceInfoRequestBody(p.DeviceInfo)
	}
	return body
}

// NewVerifyBackupCodeRequestBody builds the HTTP request body from the payload
// of the "VerifyBackupCode" endpoint of the "auth" service.
func NewVerifyBackupCodeRequestBody(p *auth.VerifyBackupCodePayload) *VerifyBackupCodeRequestBody {
	body := &VerifyBackupCodeRequestBody{
		Identifier: p.Identifier,
		BackupCode: p.BackupCode,
	}
	if p.DeviceInfo != nil {
		body.DeviceInfo = marshalAuthDeviceInfoToDeviceInfoRequestBody(p.DeviceInfo)
	}
	return body
}

// NewGenerateSecretTOTPSecretOK builds a "auth" service "GenerateSecret"
// endpoint result from a HTTP "OK" response.
func NewGenerateSecretTOTPSecretOK(body *GenerateSecretResponseBody) *auth.TOTPSecret {
	v := &auth.TOTPSecret{
		TotpURL: *body.TotpURL,
		Issuer:  *body.Issuer,
	}
	v.BackupCodes = make([]string, len(body.BackupCodes))
	for i, val := range body.BackupCodes {
		v.BackupCodes[i] = val
	}

	return v
}

// NewGenerateSecretInvalidInput builds a auth service GenerateSecret endpoint
// invalid_input error.
func NewGenerateSecretInvalidInput(body string) auth.InvalidInput {
	v := auth.InvalidInput(body)

	return v
}

// NewGenerateSecretServiceUnavailable builds a auth service GenerateSecret
// endpoint service_unavailable error.
func NewGenerateSecretServiceUnavailable(body string) auth.ServiceUnavailable {
	v := auth.ServiceUnavailable(body)

	return v
}

// NewVerifyTOTPAuthResponseOK builds a "auth" service "VerifyTOTP" endpoint
// result from a HTTP "OK" response.
func NewVerifyTOTPAuthResponseOK(body *VerifyTOTPResponseBody, sessionToken string) *auth.AuthResponse {
	v := &auth.AuthResponse{
		Success: *body.Success,
		Message: *body.Message,
	}
	v.User = unmarshalUserResponseBodyToAuthUser(body.User)
	v.SessionToken = sessionToken

	return v
}

// NewVerifyTOTPInvalidInput builds a auth service VerifyTOTP endpoint
// invalid_input error.
func NewVerifyTOTPInvalidInput(body string) auth.InvalidInput {
	v := auth.InvalidInput(body)

	return v
}

// NewVerifyTOTPInvalidOtp builds a auth service VerifyTOTP endpoint
// invalid_otp error.
func NewVerifyTOTPInvalidOtp(body string) auth.InvalidOtp {
	v := auth.InvalidOtp(body)

	return v
}

// NewVerifyTOTPUserNotFound builds a auth service VerifyTOTP endpoint
// user_not_found error.
func NewVerifyTOTPUserNotFound(body string) auth.UserNotFound {
	v := auth.UserNotFound(body)

	return v
}

// NewVerifyBackupCodeBackupCodeResponseOK builds a "auth" service
// "VerifyBackupCode" endpoint result from a HTTP "OK" response.
func NewVerifyBackupCodeBackupCodeResponseOK(body *VerifyBackupCodeResponseBody, sessionToken string) *auth.BackupCodeResponse {
	v := &auth.BackupCodeResponse{
		Success:        *body.Success,
		Message:        *body.Message,
		RemainingCodes: *body.RemainingCodes,
	}
	v.User = unmarshalUserResponseBodyToAuthUser(body.User)
	v.SessionToken = sessionToken

	return v
}

// NewVerifyBackupCodeInvalidInput builds a auth service VerifyBackupCode
// endpoint invalid_input error.
func NewVerifyBackupCodeInvalidInput(body string) auth.InvalidInput {
	v := auth.InvalidInput(body)

	return v
}

// NewVerifyBackupCodeInvalidOtp builds a auth service VerifyBackupCode
// endpoint invalid_otp error.
func NewVerifyBackupCodeInvalidOtp(body string) auth.InvalidOtp {
	v := auth.InvalidOtp(body)

	return v
}

// NewVerifyBackupCodeUserNotFound builds a auth service VerifyBackupCode
// endpoint user_not_found error.
func NewVerifyBackupCodeUserNotFound(body string) auth.UserNotFound {
	v := auth.UserNotFound(body)

	return v
}

// NewRefreshSessionSimpleResponseOK builds a "auth" service "RefreshSession"
// endpoint result from a HTTP "OK" response.
func NewRefreshSessionSimpleResponseOK(body *RefreshSessionResponseBody, sessionToken *string) *auth.SimpleResponse {
	v := &auth.SimpleResponse{
		Success: *body.Success,
		Message: *body.Message,
	}
	v.SessionToken = sessionToken

	return v
}

// NewRefreshSessionUnauthorized builds a auth service RefreshSession endpoint
// unauthorized error.
func NewRefreshSessionUnauthorized(body string) auth.Unauthorized {
	v := auth.Unauthorized(body)

	return v
}

// NewLogoutSimpleResponseOK builds a "auth" service "Logout" endpoint result
// from a HTTP "OK" response.
func NewLogoutSimpleResponseOK(body *LogoutResponseBody) *auth.SimpleResponse {
	v := &auth.SimpleResponse{
		Success:      *body.Success,
		Message:      *body.Message,
		SessionToken: body.SessionToken,
	}

	return v
}

// NewGetUserProfileUserOK builds a "auth" service "GetUserProfile" endpoint
// result from a HTTP "OK" response.
func NewGetUserProfileUserOK(body *GetUserProfileResponseBody) *auth.User {
	v := &auth.User{
		ID:         *body.ID,
		Identifier: *body.Identifier,
		CreatedAt:  *body.CreatedAt,
		LastLogin:  body.LastLogin,
		IsVerified: *body.IsVerified,
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]string, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}

	return v
}

// NewGetUserProfileUnauthorized builds a auth service GetUserProfile endpoint
// unauthorized error.
func NewGetUserProfileUnauthorized(body string) auth.Unauthorized {
	v := auth.Unauthorized(body)

	return v
}

// NewGetUserProfileUserNotFound builds a auth service GetUserProfile endpoint
// user_not_found error.
func NewGetUserProfileUserNotFound(body string) auth.UserNotFound {
	v := auth.UserNotFound(body)

	return v
}

// ValidateGenerateSecretResponseBody runs the validations defined on
// GenerateSecretResponseBody
func ValidateGenerateSecretResponseBody(body *GenerateSecretResponseBody) (err error) {
	if body.TotpURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("totp_url", "body"))
	}
	if body.BackupCodes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("backup_codes", "body"))
	}
	if body.Issuer == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("issuer", "body"))
	}
	return
}

// ValidateVerifyTOTPResponseBody runs the validations defined on
// VerifyTOTPResponseBody
func ValidateVerifyTOTPResponseBody(body *VerifyTOTPResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVerifyBackupCodeResponseBody runs the validations defined on
// VerifyBackupCodeResponseBody
func ValidateVerifyBackupCodeResponseBody(body *VerifyBackupCodeResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.RemainingCodes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("remaining_codes", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRefreshSessionResponseBody runs the validations defined on
// RefreshSessionResponseBody
func ValidateRefreshSessionResponseBody(body *RefreshSessionResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLogoutResponseBody runs the validations defined on LogoutResponseBody
func ValidateLogoutResponseBody(body *LogoutResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetUserProfileResponseBody runs the validations defined on
// GetUserProfileResponseBody
func ValidateGetUserProfileResponseBody(body *GetUserProfileResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Identifier == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("identifier", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.IsVerified == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_verified", "body"))
	}
	return
}

// ValidateUserResponseBody runs the validations defined on UserResponseBody
func ValidateUserResponseBody(body *UserResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Identifier == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("identifier", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.IsVerified == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_verified", "body"))
	}
	return
}
