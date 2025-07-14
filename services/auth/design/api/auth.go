package design

import (
	. "goa.design/goa/v3/dsl"
)

// AuthService defines the authentication microservice with OTP strategy
var _ = Service("auth", func() {
	Description("Authentication microservice with OTP support")

	// HTTP transport for client communication only
	HTTP(func() {
		Path("/api/auth")
	})

	// gRPC transport for inter-service communication
	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	// Global error definitions for the service
	Error("invalid_input", String, "Invalid input parameters")
	Error("rate_limited", String, "Too many requests")
	Error("service_unavailable", String, "Service temporarily unavailable")
	Error("invalid_otp", String, "Invalid or expired OTP")
	Error("user_not_found", String, "User not found")
	Error("unauthorized", String, "Unauthorized access")

	// Authentication methods
	// DONE in frontend
	// NOT TESTED
	Method("GenerateSecret", func() {
		Description("Generate TOTP secret for new user registration")

		Payload(func() {
			Field(1, "identifier", String, "User identifier (username/email)", func() {
				Example("user@example.com")

				MinLength(3)
				MaxLength(100)
			})

			Required("identifier")
		})

		Result(TOTPSecret)

		HTTP(func() {
			POST("/totp/generate")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// DONE in frontend
	// NOT TESTED
	Method("VerifyTOTP", func() {
		Description("Verify TOTP code and authenticate user")

		Payload(func() {
			Field(1, "identifier", String, "User identifier")
			Field(2, "totp_code", String, "6-digit TOTP code from authenticator app", func() {
				Pattern("^[0-9]{6}$")
				Example("123456")
			})
			Field(3, "device_info", DeviceInfo, "Device information for security")
			Required("identifier", "totp_code")
		})

		Result(AuthResponse)

		HTTP(func() {
			POST("/totp/verify")

			// Set authentication cookie on successful verification
			Response(StatusOK, func() {
				Cookie("session_token:session", String, "Authentication session cookie")
			})
			Response("invalid_otp", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("user_not_found", StatusNotFound)
		})
	})

	// DONE in frontend
	// NOT TESTED
	Method("VerifyBackupCode", func() {
		Description("Verify backup recovery code as alternative to TOTP")

		Payload(func() {
			Field(1, "identifier", String, "User identifier")
			Field(2, "backup_code", String, "8-character backup code", func() {
				MinLength(8)
				MaxLength(8)
				Example("ABC12345")
			})
			Field(3, "device_info", DeviceInfo, "Device information for security")
			Required("identifier", "backup_code")
		})

		Result(BackupCodeResponse)

		HTTP(func() {
			POST("/backup/verify")

			Response(StatusOK, func() {
				Cookie("session_token:session", String, "Authentication session cookie")
			})
			Response("invalid_otp", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("user_not_found", StatusNotFound)
		})
	})

	// DONE in frontend
	// NOT TESTED
	Method("RefreshSession", func() {
		Description("Refresh user session using existing token")

		Payload(func() {
			Field(1, "session_token", String, "Current session token")

			Required("session_token")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/session/refresh")
			Cookie("session_token:session")

			Response(StatusOK, func() {
				Cookie("session_token:session", String, "Refreshed session token")
			})
			Response("unauthorized", StatusUnauthorized)
		})
	})

	// Not valid for frontend, only for gRPC inter-service communication
	Method("ValidateUser", func() {
		Description("Validate user session and get user information - for gRPC inter-service communication")

		Payload(func() {
			Field(1, "session_token", String, "Session token for validation")
			Required("session_token")
		})

		Result(UserValidationResponse)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodeUnauthenticated)
		})
	})

	// Not valid for frontend, only for gRPC inter-service communication
	Method("GetUserByID", func() {
		Description("Get user information by user ID - for gRPC inter-service communication")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID to retrieve", func() {
				Minimum(1)
			})

			Required("user_id")
		})

		Result(User)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("user_not_found", CodeNotFound)
		})
	})

	// DONE in frontend
	// NOT TESTED
	Method("Logout", func() {
		Description("Logout user and invalidate session")

		Payload(func() {
			Field(1, "session_token", String, "Session token to invalidate")

			Required("session_token")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/logout")
			Cookie("session_token:session")

			Response(StatusOK)
		})
	})

	// DONE in frontend
	// NOT TESTED
	Method("GetUserProfile", func() {
		Description("Get authenticated user profile")
		Payload(func() {
			Field(1, "session_token", String, "Session token for authentication")
			Required("session_token")
		})

		Result(User)

		HTTP(func() {
			GET("/profile")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("user_not_found", StatusNotFound)
		})
	})
})
