-- name: CreateBackupCodes :copyfrom
INSERT INTO backup_codes (
    user_id,
    code_hash
) VALUES (
    $1, $2
);

-- name: GetUserBackupCodes :many
SELECT * FROM backup_codes
WHERE user_id = $1 
AND used_at IS NULL
ORDER BY created_at ASC;

-- name: GetBackupCodeByHash :one
SELECT * FROM backup_codes
WHERE user_id = $1 
AND code_hash = $2 
AND used_at IS NULL;

-- name: UseBackupCode :one
UPDATE backup_codes
SET used_at = NOW()
WHERE user_id = $1 
AND code_hash = $2 
AND used_at IS NULL
RETURNING *;

-- name: GetUsedBackupCodes :many
SELECT * FROM backup_codes
WHERE user_id = $1 
AND used_at IS NOT NULL
ORDER BY used_at DESC;

-- name: CountAvailableBackupCodes :one
SELECT COUNT(*) FROM backup_codes
WHERE user_id = $1 
AND used_at IS NULL;

-- name: DeleteUserBackupCodes :exec
DELETE FROM backup_codes
WHERE user_id = $1;

-- name: DeleteUsedBackupCodes :exec
DELETE FROM backup_codes
WHERE user_id = $1 
AND used_at IS NOT NULL;
