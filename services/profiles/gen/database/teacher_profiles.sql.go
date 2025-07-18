// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: teacher_profiles.sql

package profilesdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTeacherProfile = `-- name: CreateTeacherProfile :one
INSERT INTO teacher_profiles (
    profile_id, position
) VALUES (
    $1, $2
) RETURNING id, profile_id, position, created_at, updated_at
`

type CreateTeacherProfileParams struct {
	ProfileID int64
	Position  string
}

func (q *Queries) CreateTeacherProfile(ctx context.Context, arg CreateTeacherProfileParams) (TeacherProfile, error) {
	row := q.db.QueryRow(ctx, createTeacherProfile, arg.ProfileID, arg.Position)
	var i TeacherProfile
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCompleteTeacherProfile = `-- name: GetCompleteTeacherProfile :one
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
WHERE p.user_id = $1 AND p.is_active = true
`

type GetCompleteTeacherProfileRow struct {
	UserID    int64
	Role      string
	FirstName string
	LastName  string
	Email     string
	Phone     pgtype.Text
	AvatarUrl pgtype.Text
	Bio       pgtype.Text
	IsActive  bool
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	Position  string
}

func (q *Queries) GetCompleteTeacherProfile(ctx context.Context, userID int64) (GetCompleteTeacherProfileRow, error) {
	row := q.db.QueryRow(ctx, getCompleteTeacherProfile, userID)
	var i GetCompleteTeacherProfileRow
	err := row.Scan(
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
		&i.Position,
	)
	return i, err
}

const getTeacherProfileByProfileId = `-- name: GetTeacherProfileByProfileId :one
SELECT id, profile_id, position, created_at, updated_at FROM teacher_profiles
WHERE profile_id = $1
`

func (q *Queries) GetTeacherProfileByProfileId(ctx context.Context, profileID int64) (TeacherProfile, error) {
	row := q.db.QueryRow(ctx, getTeacherProfileByProfileId, profileID)
	var i TeacherProfile
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeacherProfileByUserId = `-- name: GetTeacherProfileByUserId :one
SELECT tp.id, tp.profile_id, tp.position, tp.created_at, tp.updated_at FROM teacher_profiles tp
JOIN profiles p ON tp.profile_id = p.id
WHERE p.user_id = $1 AND p.is_active = true
`

func (q *Queries) GetTeacherProfileByUserId(ctx context.Context, userID int64) (TeacherProfile, error) {
	row := q.db.QueryRow(ctx, getTeacherProfileByUserId, userID)
	var i TeacherProfile
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTeacherProfile = `-- name: UpdateTeacherProfile :one
UPDATE teacher_profiles 
SET 
    position = COALESCE($2, position),
    updated_at = NOW()
WHERE profile_id = $1
RETURNING id, profile_id, position, created_at, updated_at
`

type UpdateTeacherProfileParams struct {
	ProfileID int64
	Position  string
}

func (q *Queries) UpdateTeacherProfile(ctx context.Context, arg UpdateTeacherProfileParams) (TeacherProfile, error) {
	row := q.db.QueryRow(ctx, updateTeacherProfile, arg.ProfileID, arg.Position)
	var i TeacherProfile
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
