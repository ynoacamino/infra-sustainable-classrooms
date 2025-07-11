-- name: CreateStudentProfile :one
INSERT INTO student_profiles (
    profile_id, grade_level, major
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetStudentProfileByProfileId :one
SELECT * FROM student_profiles
WHERE profile_id = $1;

-- name: GetStudentProfileByUserId :one
SELECT sp.* FROM student_profiles sp
JOIN profiles p ON sp.profile_id = p.id
WHERE p.user_id = $1 AND p.is_active = true;

-- name: UpdateStudentProfile :one
UPDATE student_profiles 
SET 
    grade_level = COALESCE($2, grade_level),
    major = COALESCE($3, major),
    updated_at = NOW()
WHERE profile_id = $1
RETURNING *;

-- name: GetCompleteStudentProfile :one
SELECT 
    p.user_id,
    p.role,
    p.first_name,
    p.last_name,
    p.email,
    p.phone,
    p.avatar_url,
    p.bio,
    p.is_active,
    p.created_at,
    p.updated_at,
    sp.grade_level,
    sp.major
FROM profiles p
JOIN student_profiles sp ON p.id = sp.profile_id
WHERE p.user_id = $1 AND p.is_active = true;
