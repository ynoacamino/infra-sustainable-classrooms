package design

import (
	. "goa.design/goa/v3/dsl"
)

// User type definition
var User = Type("User", func() {
	Description("User information")

	Field(1, "id", Int64, "User unique identifier")
	Field(2, "identifier", String, "Phone number or email")
	Field(3, "created_at", Int64, "Account creation timestamp in milliseconds")
	Field(4, "last_login", Int64, "Last login timestamp in milliseconds")
	Field(5, "is_verified", Boolean, "Account verification status")
	Field(6, "metadata", MapOf(String, String), "Additional user metadata")

	Required("id", "identifier", "created_at", "is_verified")
})

// DeviceInfo type for security tracking
var DeviceInfo = Type("DeviceInfo", func() {
	Description("Device information for security purposes")

	Field(1, "user_agent", String, "Browser/app user agent")
	Field(2, "ip_address", String, "Client IP address")
	Field(3, "device_id", String, "Unique device identifier")
	Field(4, "platform", String, "Platform (web, ios, android)")
})

// Session type for session management
var Session = Type("Session", func() {
	Description("User session information")

	Field(1, "id", Int64, "Session unique identifier")
	Field(2, "user_id", Int64, "Associated user ID")
	Field(3, "created_at", Int64, "Session creation timestamp in milliseconds")
	Field(4, "expires_at", Int64, "Session expiration timestamp in milliseconds")
	Field(5, "last_accessed", Int64, "Last access timestamp in milliseconds")
	Field(6, "is_active", Boolean, "Session active status")
	Field(7, "user_agent", String, "Browser/app user agent")
	Field(8, "ip_address", String, "Client IP address")
	Field(9, "device_id", String, "Device identifier")
	Field(10, "platform", String, "Platform (web, ios, android)")

	Required("id", "user_id", "created_at", "expires_at", "is_active")
})

// BackupCode type for backup codes
var BackupCode = Type("BackupCode", func() {
	Description("Backup recovery code information")

	Field(1, "id", Int64, "Backup code unique identifier")
	Field(2, "user_id", Int64, "Associated user ID")
	Field(3, "created_at", Int64, "Creation timestamp in milliseconds")
	Field(4, "used_at", Int64, "Usage timestamp in milliseconds")

	Required("id", "user_id", "created_at")
})

// AuthAttempt type for authentication attempts
var AuthAttempt = Type("AuthAttempt", func() {
	Description("Authentication attempt information")

	Field(1, "id", Int64, "Attempt unique identifier")
	Field(2, "identifier", String, "User identifier")
	Field(3, "ip_address", String, "Client IP address")
	Field(4, "attempt_type", String, "Type of authentication attempt")
	Field(5, "success", Boolean, "Whether the attempt was successful")
	Field(6, "attempted_at", Int64, "Attempt timestamp in milliseconds")

	Required("id", "identifier", "ip_address", "attempt_type", "success", "attempted_at")
})

// TOTPSecret type for TOTP secret information
var TOTPSecret = Type("TOTPSecret", func() {
	Description("TOTP secret information for user registration")

	Field(1, "totp_url", String, "TOTP URL in otpauth:// format for authenticator apps")
	Field(2, "backup_codes", ArrayOf(String), "Backup recovery codes")
	Field(3, "issuer", String, "Service name for authenticator app")

	Required("totp_url", "backup_codes", "issuer")
})

// AuthResponse type for authentication responses
var AuthResponse = Type("AuthResponse", func() {
	Description("Standard authentication response")

	Field(1, "success", Boolean, "Authentication success status")
	Field(2, "message", String, "Response message")
	Field(3, "user", User, "User information")
	Field(4, "session_token", String, "Session token for cookie")

	Required("success", "message", "user", "session_token")
})

// BackupCodeResponse type for backup code verification
var BackupCodeResponse = Type("BackupCodeResponse", func() {
	Description("Backup code verification response")

	Field(1, "success", Boolean, "Authentication success status")
	Field(2, "message", String, "Response message")
	Field(3, "user", User, "User information")
	Field(4, "remaining_codes", Int, "Number of remaining backup codes")
	Field(5, "session_token", String, "Session token for cookie")

	Required("success", "message", "user", "remaining_codes", "session_token")
})

// ValidationResponse type for session validation
var ValidationResponse = Type("ValidationResponse", func() {
	Description("Session validation response")

	Field(1, "valid", Boolean, "Session validity status")
	Field(2, "user", User, "User information if session is valid")

	Required("valid")
})

// UserValidationResponse type for user validation with complete info
var UserValidationResponse = Type("UserValidationResponse", func() {
	Description("User validation response with complete user and session information")

	Field(1, "valid", Boolean, "Session validity status")
	Field(2, "user", User, "Complete user information if session is valid")
	Field(3, "session", Session, "Session information if valid")

	Required("valid")
})

// SimpleResponse type for simple success/message responses
var SimpleResponse = Type("SimpleResponse", func() {
	Description("Simple response with success status and message")

	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")
	Field(3, "session_token", String, "Session token for cookie (when applicable)")

	Required("success", "message")
})
