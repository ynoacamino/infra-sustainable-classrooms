package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("profiles", func() {
	Description("Profiles microservice for managing student and teacher profiles")

	// HTTP transport for client communication
	HTTP(func() {
		Path("/api/profiles")
	})

	// gRPC transport for inter-service communication
	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	// Global error definitions
	Error("invalid_input", String, "Invalid input parameters")
	Error("unauthorized", String, "Unauthorized access")
	Error("profile_not_found", String, "Profile not found")
	Error("permission_denied", String, "Permission denied")
	Error("profile_already_exists", String, "Profile already exists")
	Error("invalid_role", String, "Invalid role specified")

	// === STUDENT PROFILE CREATION ===
	// DONE in frontend
	// NOT TESTED
	Method("CreateStudentProfile", func() {
		Description("Create a new student profile with basic information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "first_name", String, "First name")
			Field(3, "last_name", String, "Last name")
			Field(4, "email", String, "Email address", func() {
				Format(FormatEmail)
			})
			Field(5, "phone", String, "Phone number")
			Field(6, "avatar_url", String, "Profile picture URL")
			Field(7, "bio", String, "Biography/description")
			Field(8, "grade_level", String, "Grade level (1-12, undergraduate, graduate)")
			Field(9, "major", String, "Major/field of study")

			Required("session_token", "first_name", "last_name", "email", "grade_level")
		})

		Result(StudentProfileResponse)

		HTTP(func() {
			POST("/student")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("profile_already_exists", StatusConflict)
		})
	})

	// === TEACHER PROFILE CREATION ===
	// DONE in frontend
	// NOT TESTED
	Method("CreateTeacherProfile", func() {
		Description("Create a new teacher profile with basic information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(1, "first_name", String, "First name")
			Field(3, "last_name", String, "Last name")
			Field(4, "email", String, "Email address", func() {
				Format(FormatEmail)
			})
			Field(5, "phone", String, "Phone number")
			Field(6, "avatar_url", String, "Profile picture URL")
			Field(7, "bio", String, "Biography/description")
			Field(8, "position", String, "Position/title")

			Required("session_token", "first_name", "last_name", "email", "position")
		})

		Result(TeacherProfileResponse)

		HTTP(func() {
			POST("/teacher")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("profile_already_exists", StatusConflict)
		})
	})

	// === PROFILE RETRIEVAL ===
	// DONE in frontend
	// NOT TESTED
	Method("GetCompleteProfile", func() {
		Description("Get user's complete profile")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")

			Required("session_token")
		})

		Result(CompleteProfileResponse)

		HTTP(func() {
			GET("/me")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("profile_not_found", StatusNotFound)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("profile_not_found", CodeNotFound)
		})
	})

	// === PUBLIC PROFILE RETRIEVAL ===
	// DONE in frontend
	// NOT TESTED
	Method("GetPublicProfileById", func() {
		Description("Get public profile information by user ID")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID to retrieve profile for")

			Required("user_id")
		})

		Result(PublicProfileResponse)

		HTTP(func() {
			GET("/public/{user_id}")
			Param("user_id", Int64, "User ID to retrieve profile for")

			Response(StatusOK)
			Response("profile_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("profile_not_found", CodeNotFound)
			Response("invalid_input", CodeInvalidArgument)
		})
	})

	// === PROFILE UPDATES ===
	// DONE in frontend
	// NOT TESTED
	Method("UpdateProfile", func() {
		Description("Update basic profile information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "first_name", String, "Updated first name")
			Field(3, "last_name", String, "Updated last name")
			Field(4, "email", String, "Updated email address", func() {
				Format(FormatEmail)
			})
			Field(5, "phone", String, "Updated phone number")
			Field(6, "avatar_url", String, "Updated profile picture URL")
			Field(7, "bio", String, "Updated biography")
			Field(8, "major", String, "Updated major (students)")
			Field(9, "position", String, "Updated position (teachers)")

			Required("session_token")
		})

		Result(ProfileResponse)

		HTTP(func() {
			PUT("/me")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("profile_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// === INTER-SERVICE COMMUNICATION (gRPC only) ===
	// Not used in frontend, only for inter-service communication
	Method("ValidateUserRole", func() {
		Description("Validate user role for inter-service communication")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID to validate")

			Required("user_id")
		})

		Result(RoleValidationResponse)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("profile_not_found", CodeNotFound)
			Response("permission_denied", CodePermissionDenied)
		})
	})
})
