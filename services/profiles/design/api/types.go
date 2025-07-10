package design

import (
	. "goa.design/goa/v3/dsl"
)

// === BASIC RESPONSE TYPES ===
var SimpleResponse = Type("SimpleResponse", func() {
	Description("Basic response with success status and message")

	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")

	Required("success", "message")
})

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
	Field(9, "created_at", String, "Profile creation timestamp")
	Field(10, "updated_at", String, "Last update timestamp")
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
	Field(7, "school_name", String, "School/University name")
	Field(8, "department", String, "Department (for teachers)")
	Field(9, "subjects", ArrayOf(String), "Subjects/courses")
	Field(10, "is_active", Boolean, "Whether profile is active")

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
	Field(9, "school_name", String, "School/University name")
	Field(10, "major", String, "Major/field of study")
	Field(11, "subjects", ArrayOf(String), "Enrolled subjects/courses")
	Field(12, "created_at", String, "Profile creation timestamp")
	Field(13, "updated_at", String, "Last update timestamp")
	Field(14, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "first_name", "last_name", "email", "grade_level", "school_name", "created_at", "is_active")
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
	Field(8, "department", String, "Department name")
	Field(9, "position", String, "Position/title")
	Field(10, "subjects_taught", ArrayOf(String), "Subjects/courses taught")
	Field(11, "school_name", String, "School/University name")
	Field(12, "created_at", String, "Profile creation timestamp")
	Field(13, "updated_at", String, "Last update timestamp")
	Field(14, "is_active", Boolean, "Whether profile is active")

	Required("user_id", "first_name", "last_name", "email", "department", "position", "subjects_taught", "created_at", "is_active")
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
	Field(9, "created_at", String, "Profile creation timestamp")
	Field(10, "updated_at", String, "Last update timestamp")
	Field(11, "is_active", Boolean, "Whether profile is active")

	// Role-specific fields
	Field(12, "grade_level", String, "Grade level (for students)")
	Field(13, "school_name", String, "School/University name")
	Field(14, "major", String, "Major/field of study (for students)")
	Field(15, "subjects", ArrayOf(String), "Enrolled subjects (for students)")
	Field(16, "department", String, "Department (for teachers)")
	Field(17, "position", String, "Position/title (for teachers)")
	Field(18, "subjects_taught", ArrayOf(String), "Subjects taught (for teachers)")

	Required("user_id", "role", "first_name", "last_name", "email", "created_at", "is_active")
})

// === INTER-SERVICE COMMUNICATION TYPES ===
var RoleValidationResponse = Type("RoleValidationResponse", func() {
	Description("Response for role validation")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role")
	Field(3, "has_permission", Boolean, "Whether user has required permission")
	Field(4, "permissions", ArrayOf(String), "User permissions")
	Field(5, "department", String, "Department (if applicable)")
	Field(6, "subjects", ArrayOf(String), "Associated subjects")

	Required("user_id", "role", "has_permission")
})

var BasicUserInfo = Type("BasicUserInfo", func() {
	Description("Basic user information for inter-service communication")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "role", String, "User role")
	Field(3, "first_name", String, "First name")
	Field(4, "last_name", String, "Last name")
	Field(5, "email", String, "Email address")
	Field(6, "is_active", Boolean, "Whether profile is active")
	Field(7, "school_name", String, "School/University name")
	Field(8, "department", String, "Department (for teachers)")
	Field(9, "grade_level", String, "Grade level (for students)")

	Required("user_id", "role", "first_name", "last_name", "email", "is_active")
})
