// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: profiles.sql

package profilesdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkProfileExists = `-- name: CheckProfileExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE user_id = $1 AND is_active = true)
`

func (q *Queries) CheckProfileExists(ctx context.Context, userID int64) (bool, error) {
	row := q.db.QueryRow(ctx, checkProfileExists, userID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createProfile = `-- name: CreateProfile :one
INSERT INTO profiles (
    user_id, role, first_name, last_name, email, phone, avatar_url, bio
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, user_id, role, first_name, last_name, email, phone, avatar_url, bio, is_active, created_at, updated_at
`

type CreateProfileParams struct {
	UserID    int64
	Role      string
	FirstName string
	LastName  string
	Email     string
	Phone     pgtype.Text
	AvatarUrl pgtype.Text
	Bio       pgtype.Text
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, createProfile,
		arg.UserID,
		arg.Role,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.AvatarUrl,
		arg.Bio,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.AvatarUrl,
		&i.Bio,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deactivateProfile = `-- name: DeactivateProfile :exec
UPDATE profiles 
SET is_active = false, updated_at = NOW()
WHERE user_id = $1
`

func (q *Queries) DeactivateProfile(ctx context.Context, userID int64) error {
	_, err := q.db.Exec(ctx, deactivateProfile, userID)
	return err
}

const getProfileByEmail = `-- name: GetProfileByEmail :one
SELECT id, user_id, role, first_name, last_name, email, phone, avatar_url, bio, is_active, created_at, updated_at FROM profiles
WHERE email = $1 AND is_active = true
`

func (q *Queries) GetProfileByEmail(ctx context.Context, email string) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfileByEmail, email)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.AvatarUrl,
		&i.Bio,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProfileByUserId = `-- name: GetProfileByUserId :one
SELECT id, user_id, role, first_name, last_name, email, phone, avatar_url, bio, is_active, created_at, updated_at FROM profiles
WHERE user_id = $1 AND is_active = true
`

func (q *Queries) GetProfileByUserId(ctx context.Context, userID int64) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfileByUserId, userID)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.AvatarUrl,
		&i.Bio,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProfileRole = `-- name: GetProfileRole :one
SELECT role FROM profiles WHERE user_id = $1 AND is_active = true
`

func (q *Queries) GetProfileRole(ctx context.Context, userID int64) (string, error) {
	row := q.db.QueryRow(ctx, getProfileRole, userID)
	var role string
	err := row.Scan(&role)
	return role, err
}

const getPublicProfileByUserId = `-- name: GetPublicProfileByUserId :one
SELECT user_id, role, first_name, last_name, avatar_url, bio, is_active
FROM profiles
WHERE user_id = $1 AND is_active = true
`

type GetPublicProfileByUserIdRow struct {
	UserID    int64
	Role      string
	FirstName string
	LastName  string
	AvatarUrl pgtype.Text
	Bio       pgtype.Text
	IsActive  bool
}

func (q *Queries) GetPublicProfileByUserId(ctx context.Context, userID int64) (GetPublicProfileByUserIdRow, error) {
	row := q.db.QueryRow(ctx, getPublicProfileByUserId, userID)
	var i GetPublicProfileByUserIdRow
	err := row.Scan(
		&i.UserID,
		&i.Role,
		&i.FirstName,
		&i.LastName,
		&i.AvatarUrl,
		&i.Bio,
		&i.IsActive,
	)
	return i, err
}

const updateProfile = `-- name: UpdateProfile :one
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
RETURNING id, user_id, role, first_name, last_name, email, phone, avatar_url, bio, is_active, created_at, updated_at
`

type UpdateProfileParams struct {
	UserID    int64
	FirstName string
	LastName  string
	Email     string
	Phone     pgtype.Text
	AvatarUrl pgtype.Text
	Bio       pgtype.Text
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, updateProfile,
		arg.UserID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.AvatarUrl,
		arg.Bio,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.AvatarUrl,
		&i.Bio,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
