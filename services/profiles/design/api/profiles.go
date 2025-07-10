package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("profiles", func() {
	Description("Profiles microservice for managing student and teacher profiles")

	// HTTP transport for client communication
	HTTP(func() {
		Path("/profiles")
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
	Method("CreateStudentProfile", func() {
		Description("Create a new student profile with basic information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User ID from auth service")
			Field(3, "first_name", String, "First name")
			Field(4, "last_name", String, "Last name")
			Field(5, "email", String, "Email address", func() {
				Format(FormatEmail)
			})
			Field(6, "phone", String, "Phone number")
			Field(7, "avatar_url", String, "Profile picture URL")
			Field(8, "bio", String, "Biography/description")
			Field(9, "grade_level", String, "Grade level (1-12, undergraduate, graduate)")
			Field(10, "school_name", String, "School/University name")
			Field(11, "major", String, "Major/field of study")
			Field(12, "subjects", ArrayOf(String), "Enrolled subjects/courses")

			Required("session_token", "user_id", "first_name", "last_name", "email", "grade_level", "school_name")
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
	Method("CreateTeacherProfile", func() {
		Description("Create a new teacher profile with basic information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User ID from auth service")
			Field(3, "first_name", String, "First name")
			Field(4, "last_name", String, "Last name")
			Field(5, "email", String, "Email address", func() {
				Format(FormatEmail)
			})
			Field(6, "phone", String, "Phone number")
			Field(7, "avatar_url", String, "Profile picture URL")
			Field(8, "bio", String, "Biography/description")
			Field(9, "department", String, "Department name")
			Field(10, "position", String, "Position/title")
			Field(11, "subjects_taught", ArrayOf(String), "Subjects/courses taught")
			Field(12, "school_name", String, "School/University name")

			Required("session_token", "user_id", "first_name", "last_name", "email", "department", "position", "subjects_taught")
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
	Method("GetMyProfile", func() {
		Description("Get current user's complete profile")

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
	})

	// === PROFILE UPDATES ===
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
			Field(9, "subjects", ArrayOf(String), "Updated subjects (students)")
			Field(10, "department", String, "Updated department (teachers)")
			Field(11, "position", String, "Updated position (teachers)")
			Field(12, "subjects_taught", ArrayOf(String), "Updated subjects taught (teachers)")

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
	Method("ValidateUserRole", func() {
		Description("Validate user role for inter-service communication")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID to validate")
			Field(2, "required_role", String, "Required role")

			Required("user_id", "required_role")
		})

		Result(RoleValidationResponse)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("profile_not_found", CodeNotFound)
			Response("permission_denied", CodePermissionDenied)
		})
	})

	Method("GetUserBasicInfo", func() {
		Description("Get basic user information for inter-service communication")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID")

			Required("user_id")
		})

		Result(BasicUserInfo)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("profile_not_found", CodeNotFound)
		})
	})
})
