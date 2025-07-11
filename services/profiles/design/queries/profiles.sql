-- name: CreateProfile :one
INSERT INTO profiles (
    user_id, role, first_name, last_name, email, phone, avatar_url, bio
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetProfileByUserId :one
SELECT * FROM profiles
WHERE user_id = $1 AND is_active = true;

-- name: GetProfileByEmail :one
SELECT * FROM profiles
WHERE email = $1 AND is_active = true;

-- name: GetPublicProfileByUserId :one
SELECT user_id, role, first_name, last_name, avatar_url, bio, is_active
FROM profiles
WHERE user_id = $1 AND is_active = true;

-- name: UpdateProfile :one
UPDATE profiles 
SET 
    first_name = COALESCE($2, first_name),
    last_name = COALESCE($3, last_name),
    email = COALESCE($4, email),
    phone = COALESCE($5, phone),
    avatar_url = COALESCE($6, avatar_url),
    bio = COALESCE($7, bio),
    updated_at = NOW()
WHERE user_id = $1 AND is_active = true
RETURNING *;

-- name: DeactivateProfile :exec
UPDATE profiles 
SET is_active = false, updated_at = NOW()
WHERE user_id = $1;

-- name: CheckProfileExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE user_id = $1 AND is_active = true);

-- name: GetProfileRole :one
SELECT role FROM profiles WHERE user_id = $1 AND is_active = true;
