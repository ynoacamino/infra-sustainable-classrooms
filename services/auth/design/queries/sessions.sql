-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    session_token,
    expires_at,
    user_agent,
    ip_address,
    device_id,
    platform
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSessionByToken :one
SELECT s.*, u.identifier, u.is_verified
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.session_token = $1 
AND s.is_active = true 
AND s.expires_at > NOW();

-- name: GetUserSessions :many
SELECT * FROM sessions
WHERE user_id = $1 
AND is_active = true 
AND expires_at > NOW()
ORDER BY last_accessed DESC;

-- name: UpdateSessionAccess :exec
UPDATE sessions
SET last_accessed = NOW()
WHERE session_token = $1 
AND is_active = true;

-- name: RefreshSession :one
UPDATE sessions
SET 
    expires_at = $2,
    last_accessed = NOW()
WHERE session_token = $1 
AND is_active = true
RETURNING *;

-- name: DeactivateSession :exec
UPDATE sessions
SET is_active = false
WHERE session_token = $1;

-- name: DeactivateUserSessions :exec
UPDATE sessions
SET is_active = false
WHERE user_id = $1 
AND is_active = true;

-- name: DeactivateAllUserSessionsExcept :exec
UPDATE sessions
SET is_active = false
WHERE user_id = $1 
AND session_token != $2 
AND is_active = true;

-- name: CleanupExpiredSessions :exec
UPDATE sessions 
SET is_active = false 
WHERE expires_at < NOW() 
AND is_active = true;

-- name: GetSessionStats :one
SELECT 
    COUNT(*) as total_sessions,
    COUNT(*) FILTER (WHERE is_active = true) as active_sessions,
    COUNT(*) FILTER (WHERE created_at > NOW() - INTERVAL '24 hours') as sessions_last_24h
FROM sessions;
