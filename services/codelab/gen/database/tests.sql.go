// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: tests.sql

package codelabdb

import (
	"context"
)

const createTest = `-- name: CreateTest :exec
INSERT INTO tests (input, output, public, exercise_id)
VALUES ($1, $2, $3, $4)
`

type CreateTestParams struct {
	Input      string
	Output     string
	Public     bool
	ExerciseID int64
}

func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) error {
	_, err := q.db.Exec(ctx, createTest,
		arg.Input,
		arg.Output,
		arg.Public,
		arg.ExerciseID,
	)
	return err
}

const deleteTest = `-- name: DeleteTest :exec
DELETE FROM tests WHERE id = $1
`

func (q *Queries) DeleteTest(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTest, id)
	return err
}

const deleteTestsByExercise = `-- name: DeleteTestsByExercise :exec
DELETE FROM tests WHERE exercise_id = $1
`

func (q *Queries) DeleteTestsByExercise(ctx context.Context, exerciseID int64) error {
	_, err := q.db.Exec(ctx, deleteTestsByExercise, exerciseID)
	return err
}

const getHiddenTestsByExercise = `-- name: GetHiddenTestsByExercise :many
SELECT id, input, output, public, exercise_id, created_at, updated_at FROM tests 
WHERE exercise_id = $1 AND public = false 
ORDER BY created_at
`

func (q *Queries) GetHiddenTestsByExercise(ctx context.Context, exerciseID int64) ([]Test, error) {
	rows, err := q.db.Query(ctx, getHiddenTestsByExercise, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Test
	for rows.Next() {
		var i Test
		if err := rows.Scan(
			&i.ID,
			&i.Input,
			&i.Output,
			&i.Public,
			&i.ExerciseID,
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

const getPublicTestsByExercise = `-- name: GetPublicTestsByExercise :many
SELECT id, input, output, public, exercise_id, created_at, updated_at FROM tests 
WHERE exercise_id = $1 AND public = true 
ORDER BY created_at
`

func (q *Queries) GetPublicTestsByExercise(ctx context.Context, exerciseID int64) ([]Test, error) {
	rows, err := q.db.Query(ctx, getPublicTestsByExercise, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Test
	for rows.Next() {
		var i Test
		if err := rows.Scan(
			&i.ID,
			&i.Input,
			&i.Output,
			&i.Public,
			&i.ExerciseID,
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

const getTestsByExercise = `-- name: GetTestsByExercise :many
SELECT id, input, output, public, exercise_id, created_at, updated_at FROM tests 
WHERE exercise_id = $1 
ORDER BY created_at
`

func (q *Queries) GetTestsByExercise(ctx context.Context, exerciseID int64) ([]Test, error) {
	rows, err := q.db.Query(ctx, getTestsByExercise, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Test
	for rows.Next() {
		var i Test
		if err := rows.Scan(
			&i.ID,
			&i.Input,
			&i.Output,
			&i.Public,
			&i.ExerciseID,
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

const updateTest = `-- name: UpdateTest :exec
UPDATE tests SET
    input = $2,
    output = $3,
    public = $4,
    updated_at = NOW()
WHERE id = $1
`

type UpdateTestParams struct {
	ID     int64
	Input  string
	Output string
	Public bool
}

func (q *Queries) UpdateTest(ctx context.Context, arg UpdateTestParams) error {
	_, err := q.db.Exec(ctx, updateTest,
		arg.ID,
		arg.Input,
		arg.Output,
		arg.Public,
	)
	return err
}
