// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: exercises.sql

package codelabdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExercise = `-- name: CreateExercise :exec
INSERT INTO exercises (
    title, description, initial_code, solution, difficulty, created_by
) VALUES (
    $1, $2, $3, $4, $5, $6
)
`

type CreateExerciseParams struct {
	Title       string
	Description string
	InitialCode string
	Solution    string
	Difficulty  string
	CreatedBy   int64
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) error {
	_, err := q.db.Exec(ctx, createExercise,
		arg.Title,
		arg.Description,
		arg.InitialCode,
		arg.Solution,
		arg.Difficulty,
		arg.CreatedBy,
	)
	return err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM exercises WHERE id = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteExercise, id)
	return err
}

const getExerciseById = `-- name: GetExerciseById :one
SELECT id, title, description, initial_code, solution, difficulty, created_by, created_at, updated_at FROM exercises WHERE id = $1
`

func (q *Queries) GetExerciseById(ctx context.Context, id int64) (Exercise, error) {
	row := q.db.QueryRow(ctx, getExerciseById, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.InitialCode,
		&i.Solution,
		&i.Difficulty,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getExerciseToResolveById = `-- name: GetExerciseToResolveById :one
SELECT 
    id,
    title,
    description,
    initial_code,
    difficulty,
    created_by,
    created_at,
    updated_at
FROM exercises WHERE id = $1
`

type GetExerciseToResolveByIdRow struct {
	ID          int64
	Title       string
	Description string
	InitialCode string
	Difficulty  string
	CreatedBy   int64
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) GetExerciseToResolveById(ctx context.Context, id int64) (GetExerciseToResolveByIdRow, error) {
	row := q.db.QueryRow(ctx, getExerciseToResolveById, id)
	var i GetExerciseToResolveByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.InitialCode,
		&i.Difficulty,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listExercises = `-- name: ListExercises :many
SELECT id, title, description, initial_code, solution, difficulty, created_by, created_at, updated_at FROM exercises ORDER BY created_at DESC
`

func (q *Queries) ListExercises(ctx context.Context) ([]Exercise, error) {
	rows, err := q.db.Query(ctx, listExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.InitialCode,
			&i.Solution,
			&i.Difficulty,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExercisesToResolve = `-- name: ListExercisesToResolve :many
SELECT
    id,
    title,
    description,
    initial_code,
    difficulty,
    created_by,
    created_at,
    updated_at
FROM exercises
ORDER BY created_at DESC
`

type ListExercisesToResolveRow struct {
	ID          int64
	Title       string
	Description string
	InitialCode string
	Difficulty  string
	CreatedBy   int64
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) ListExercisesToResolve(ctx context.Context) ([]ListExercisesToResolveRow, error) {
	rows, err := q.db.Query(ctx, listExercisesToResolve)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListExercisesToResolveRow
	for rows.Next() {
		var i ListExercisesToResolveRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.InitialCode,
			&i.Difficulty,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExercise = `-- name: UpdateExercise :exec
UPDATE exercises SET
    title = $2,
    description = $3,
    initial_code = $4,
    solution = $5,
    difficulty = $6,
    updated_at = NOW()
WHERE id = $1
`

type UpdateExerciseParams struct {
	ID          int64
	Title       string
	Description string
	InitialCode string
	Solution    string
	Difficulty  string
}

func (q *Queries) UpdateExercise(ctx context.Context, arg UpdateExerciseParams) error {
	_, err := q.db.Exec(ctx, updateExercise,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.InitialCode,
		arg.Solution,
		arg.Difficulty,
	)
	return err
}
