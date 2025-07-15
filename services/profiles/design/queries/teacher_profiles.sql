-- name: CreateTeacherProfile :one
INSERT INTO teacher_profiles (
    profile_id, position
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetTeacherProfileByProfileId :one
SELECT * FROM teacher_profiles
WHERE profile_id = $1;

-- name: GetTeacherProfileByUserId :one
SELECT tp.* FROM teacher_profiles tp
JOIN profiles p ON tp.profile_id = p.id
WHERE p.user_id = $1 AND p.is_active = true;

-- name: UpdateTeacherProfile :one
UPDATE teacher_profiles 
SET 
    position = COALESCE($2, position),
    updated_at = NOW()
WHERE profile_id = $1
RETURNING *;

-- name: GetCompleteTeacherProfile :one
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
    tp.position
FROM profiles p
JOIN teacher_profiles tp ON p.id = tp.profile_id
WHERE p.user_id = $1 AND p.is_active = true;
