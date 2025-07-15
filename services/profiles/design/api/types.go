package design

import (
	. "goa.design/goa/v3/dsl"
)

// === PROFILE TYPES ===
var ProfileResponse = Type("ProfileResponse", func() {
	Description("Basic profile information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role (student, teacher)")
	Field(3, "first_name", String, "First name")
	Field(4, "last_name", String, "Last name")
	Field(5, "email", String, "Email address")
	Field(6, "phone", String, "Phone number")
	Field(7, "avatar_url", String, "Profile picture URL")
	Field(8, "bio", String, "Biography/description")
	Field(9, "created_at", Int64, "Profile creation timestamp")
	Field(10, "updated_at", Int64, "Last update timestamp")
	Field(11, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "role", "first_name", "last_name", "email", "created_at", "is_active")
})

var PublicProfileResponse = Type("PublicProfileResponse", func() {
	Description("Public profile information (limited data)")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role")
	Field(3, "first_name", String, "First name")
	Field(4, "last_name", String, "Last name")
	Field(5, "avatar_url", String, "Profile picture URL")
	Field(6, "bio", String, "Public biography")
	Field(7, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "role", "first_name", "last_name", "is_active")
})

// === STUDENT SPECIFIC TYPES ===
var StudentProfileResponse = Type("StudentProfileResponse", func() {
	Description("Student profile information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "first_name", String, "First name")
	Field(3, "last_name", String, "Last name")
	Field(4, "email", String, "Email address")
	Field(5, "phone", String, "Phone number")
	Field(6, "avatar_url", String, "Profile picture URL")
	Field(7, "bio", String, "Biography/description")
	Field(8, "grade_level", String, "Grade level")
	Field(9, "major", String, "Major/field of study")
	Field(11, "created_at", Int64, "Profile creation timestamp")
	Field(12, "updated_at", Int64, "Last update timestamp")
	Field(13, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "first_name", "last_name", "email", "grade_level", "created_at", "is_active")
})

// === TEACHER SPECIFIC TYPES ===
var TeacherProfileResponse = Type("TeacherProfileResponse", func() {
	Description("Teacher profile information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "first_name", String, "First name")
	Field(3, "last_name", String, "Last name")
	Field(4, "email", String, "Email address")
	Field(5, "phone", String, "Phone number")
	Field(6, "avatar_url", String, "Profile picture URL")
	Field(7, "bio", String, "Biography/description")
	Field(9, "position", String, "Position/title")
	Field(10, "created_at", Int64, "Profile creation timestamp")
	Field(11, "updated_at", Int64, "Last update timestamp")
	Field(12, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "first_name", "last_name", "email", "position", "created_at", "is_active")
})

// === COMPLETE PROFILE TYPES ===
var CompleteProfileResponse = Type("CompleteProfileResponse", func() {
	Description("Complete user profile with role-specific information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role (student, teacher)")
	Field(3, "first_name", String, "First name")
	Field(4, "last_name", String, "Last name")
	Field(5, "email", String, "Email address")
	Field(6, "phone", String, "Phone number")
	Field(7, "avatar_url", String, "Profile picture URL")
	Field(8, "bio", String, "Biography/description")
	Field(9, "created_at", Int64, "Profile creation timestamp")
	Field(10, "updated_at", Int64, "Last update timestamp")
	Field(11, "is_active", Boolean, "Whether profile is active")

	// Role-specific fields
	Field(12, "grade_level", String, "Grade level (for students)")
	Field(13, "major", String, "Major/field of study (for students)")

	Field(14, "position", String, "Position/title (for teachers)")

	Required("user_id", "role", "first_name", "last_name", "email", "created_at", "is_active")
})

// === INTER-SERVICE COMMUNICATION TYPES ===
var RoleValidationResponse = Type("RoleValidationResponse", func() {
	Description("Response for role validation")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role")

	Required("user_id", "role")
})
