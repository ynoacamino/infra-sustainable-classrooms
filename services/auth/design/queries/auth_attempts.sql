-- name: RecordAuthAttempt :one
INSERT INTO auth_attempts (
    identifier,
    ip_address,
    attempt_type,
    success
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CountRecentFailedAttempts :one
SELECT COUNT(*) FROM auth_attempts
WHERE identifier = $1 
AND success = false 
AND attempted_at > NOW() - INTERVAL '15 minutes';

-- name: CountRecentFailedAttemptsByIP :one
SELECT COUNT(*) FROM auth_attempts
WHERE ip_address = $1 
AND success = false 
AND attempted_at > NOW() - INTERVAL '15 minutes';

-- name: GetLastSuccessfulAttempt :one
SELECT * FROM auth_attempts
WHERE identifier = $1 
AND success = true 
ORDER BY attempted_at DESC 
LIMIT 1;

-- name: GetRecentAttempts :many
SELECT * FROM auth_attempts
WHERE identifier = $1 
AND attempted_at > NOW() - INTERVAL '1 hour'
ORDER BY attempted_at DESC;

-- name: GetRecentAttemptsByIP :many
SELECT * FROM auth_attempts
WHERE ip_address = $1 
AND attempted_at > NOW() - INTERVAL '1 hour'
ORDER BY attempted_at DESC;

-- name: CleanupOldAuthAttempts :exec
DELETE FROM auth_attempts
WHERE attempted_at < NOW() - INTERVAL '7 days';

-- name: GetAuthStats :one
SELECT 
    COUNT(*) as total_attempts,
    COUNT(*) FILTER (WHERE success = true) as successful_attempts,
    COUNT(*) FILTER (WHERE success = false) as failed_attempts,
    COUNT(*) FILTER (WHERE attempted_at > NOW() - INTERVAL '24 hours') as attempts_last_24h,
    COUNT(*) FILTER (WHERE success = true AND attempted_at > NOW() - INTERVAL '24 hours') as successful_last_24h
FROM auth_attempts;
