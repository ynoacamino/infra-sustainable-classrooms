// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: answers.sql

package codelabdb

import (
	"context"
)

const checkIfAnswerExists = `-- name: CheckIfAnswerExists :one
SELECT 1 FROM answers
WHERE exercise_id = $1 AND user_id = $2
`

type CheckIfAnswerExistsParams struct {
	ExerciseID int64
	UserID     int64
}

func (q *Queries) CheckIfAnswerExists(ctx context.Context, arg CheckIfAnswerExistsParams) (int32, error) {
	row := q.db.QueryRow(ctx, checkIfAnswerExists, arg.ExerciseID, arg.UserID)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const countCompletedAnswersByExercise = `-- name: CountCompletedAnswersByExercise :one
SELECT COUNT(*) FROM answers 
WHERE exercise_id = $1 AND completed = true
`

func (q *Queries) CountCompletedAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countCompletedAnswersByExercise, exerciseID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countTotalAnswersByExercise = `-- name: CountTotalAnswersByExercise :one
SELECT COUNT(*) FROM answers 
WHERE exercise_id = $1
`

func (q *Queries) CountTotalAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countTotalAnswersByExercise, exerciseID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAnswer = `-- name: CreateAnswer :exec
INSERT INTO answers (exercise_id, user_id, completed)
VALUES ($1, $2, $3)
`

type CreateAnswerParams struct {
	ExerciseID int64
	UserID     int64
	Completed  bool
}

func (q *Queries) CreateAnswer(ctx context.Context, arg CreateAnswerParams) error {
	_, err := q.db.Exec(ctx, createAnswer, arg.ExerciseID, arg.UserID, arg.Completed)
	return err
}

const getAnswerByUserAndExercise = `-- name: GetAnswerByUserAndExercise :one
SELECT id, exercise_id, user_id, completed, created_at, updated_at FROM answers 
WHERE exercise_id = $1 AND user_id = $2
`

type GetAnswerByUserAndExerciseParams struct {
	ExerciseID int64
	UserID     int64
}

func (q *Queries) GetAnswerByUserAndExercise(ctx context.Context, arg GetAnswerByUserAndExerciseParams) (Answer, error) {
	row := q.db.QueryRow(ctx, getAnswerByUserAndExercise, arg.ExerciseID, arg.UserID)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.ExerciseID,
		&i.UserID,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAnswersByExercise = `-- name: ListAnswersByExercise :many
SELECT id, exercise_id, user_id, completed, created_at, updated_at FROM answers 
WHERE exercise_id = $1
ORDER BY updated_at DESC
`

func (q *Queries) ListAnswersByExercise(ctx context.Context, exerciseID int64) ([]Answer, error) {
	rows, err := q.db.Query(ctx, listAnswersByExercise, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Answer
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.ExerciseID,
			&i.UserID,
			&i.Completed,
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

const listAnswersByUser = `-- name: ListAnswersByUser :many
SELECT id, exercise_id, user_id, completed, created_at, updated_at FROM answers 
WHERE user_id = $1 
ORDER BY updated_at DESC
`

func (q *Queries) ListAnswersByUser(ctx context.Context, userID int64) ([]Answer, error) {
	rows, err := q.db.Query(ctx, listAnswersByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Answer
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.ExerciseID,
			&i.UserID,
			&i.Completed,
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

const updateAnswerCompleted = `-- name: UpdateAnswerCompleted :exec
UPDATE answers SET
    completed = $3,
    updated_at = NOW()
WHERE exercise_id = $1 AND user_id = $2
`

type UpdateAnswerCompletedParams struct {
	ExerciseID int64
	UserID     int64
	Completed  bool
}

func (q *Queries) UpdateAnswerCompleted(ctx context.Context, arg UpdateAnswerCompletedParams) error {
	_, err := q.db.Exec(ctx, updateAnswerCompleted, arg.ExerciseID, arg.UserID, arg.Completed)
	return err
}
