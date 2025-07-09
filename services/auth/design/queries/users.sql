-- name: CreateUser :one
INSERT INTO users (
    identifier,
    totp_secret,
    is_verified,
    metadata
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByIdentifier :one
SELECT * FROM users
WHERE identifier = $1;

-- name: UpdateUserTOTPSecret :one
UPDATE users
SET 
    totp_secret = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: VerifyUser :one
UPDATE users
SET 
    is_verified = true,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET 
    last_login = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserMetadata :one
UPDATE users
SET 
    metadata = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserStats :one
SELECT 
    COUNT(*) as total_users,
    COUNT(*) FILTER (WHERE is_verified = true) as verified_users,
    COUNT(*) FILTER (WHERE created_at > NOW() - INTERVAL '24 hours') as users_last_24h
FROM users;
