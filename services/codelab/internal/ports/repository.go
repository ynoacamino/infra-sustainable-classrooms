package ports

import (
	"context"

	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

// CodeLabRepository combines all repository interfaces for the codelab service
type CodeLabRepository interface {
	ExercisesRepository
	TestsRepository
	AnswersRepository
	AttemptsRepository
}

// Repositories struct holds all repository implementations
type Repositories struct {
	Exercises ExercisesRepository
	Tests     TestsRepository
	Answers   AnswersRepository
	Attempts  AttemptsRepository
}

// DatabaseRepository implements CodeLabRepository using the SQLC generated queries
type DatabaseRepository struct {
	*codelabdb.Queries
}

// NewDatabaseRepository creates a new instance of DatabaseRepository
func NewDatabaseRepository(queries *codelabdb.Queries) *DatabaseRepository {
	return &DatabaseRepository{
		Queries: queries,
	}
}

// Ensure DatabaseRepository implements all repository interfaces
var _ ExercisesRepository = (*DatabaseRepository)(nil)
var _ TestsRepository = (*DatabaseRepository)(nil)
var _ AnswersRepository = (*DatabaseRepository)(nil)
var _ AttemptsRepository = (*DatabaseRepository)(nil)
var _ CodeLabRepository = (*DatabaseRepository)(nil)

// Exercises Repository methods
func (r *DatabaseRepository) CreateExercise(ctx context.Context, arg codelabdb.CreateExerciseParams) error {
	return r.Queries.CreateExercise(ctx, arg)
}

func (r *DatabaseRepository) GetExerciseById(ctx context.Context, id int64) (codelabdb.Exercise, error) {
	return r.Queries.GetExerciseById(ctx, id)
}

func (r *DatabaseRepository) GetExerciseToResolveById(ctx context.Context, id int64) (codelabdb.GetExerciseToResolveByIdRow, error) {
	return r.Queries.GetExerciseToResolveById(ctx, id)
}

func (r *DatabaseRepository) ListExercises(ctx context.Context) ([]codelabdb.Exercise, error) {
	return r.Queries.ListExercises(ctx)
}

func (r *DatabaseRepository) ListExercisesToResolve(ctx context.Context) ([]codelabdb.ListExercisesToResolveRow, error) {
	return r.Queries.ListExercisesToResolve(ctx)
}

func (r *DatabaseRepository) UpdateExercise(ctx context.Context, arg codelabdb.UpdateExerciseParams) error {
	return r.Queries.UpdateExercise(ctx, arg)
}

func (r *DatabaseRepository) DeleteExercise(ctx context.Context, id int64) error {
	return r.Queries.DeleteExercise(ctx, id)
}

// Tests Repository methods
func (r *DatabaseRepository) CreateTest(ctx context.Context, arg codelabdb.CreateTestParams) error {
	return r.Queries.CreateTest(ctx, arg)
}

func (r *DatabaseRepository) GetTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	return r.Queries.GetTestsByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) GetPublicTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	return r.Queries.GetPublicTestsByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) GetHiddenTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	return r.Queries.GetHiddenTestsByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) UpdateTest(ctx context.Context, arg codelabdb.UpdateTestParams) error {
	return r.Queries.UpdateTest(ctx, arg)
}

func (r *DatabaseRepository) DeleteTest(ctx context.Context, id int64) error {
	return r.Queries.DeleteTest(ctx, id)
}

func (r *DatabaseRepository) DeleteTestsByExercise(ctx context.Context, exerciseID int64) error {
	return r.Queries.DeleteTestsByExercise(ctx, exerciseID)
}

// Answers Repository methods
func (r *DatabaseRepository) CheckIfAnswerExists(ctx context.Context, arg codelabdb.CheckIfAnswerExistsParams) (int32, error) {
	return r.Queries.CheckIfAnswerExists(ctx, arg)
}

func (r *DatabaseRepository) CreateAnswer(ctx context.Context, arg codelabdb.CreateAnswerParams) error {
	return r.Queries.CreateAnswer(ctx, arg)
}

func (r *DatabaseRepository) GetAnswerByUserAndExercise(ctx context.Context, arg codelabdb.GetAnswerByUserAndExerciseParams) (codelabdb.Answer, error) {
	return r.Queries.GetAnswerByUserAndExercise(ctx, arg)
}

func (r *DatabaseRepository) ListAnswersByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Answer, error) {
	return r.Queries.ListAnswersByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) ListAnswersByUser(ctx context.Context, userID int64) ([]codelabdb.Answer, error) {
	return r.Queries.ListAnswersByUser(ctx, userID)
}

func (r *DatabaseRepository) UpdateAnswerCompleted(ctx context.Context, arg codelabdb.UpdateAnswerCompletedParams) error {
	return r.Queries.UpdateAnswerCompleted(ctx, arg)
}

func (r *DatabaseRepository) CountCompletedAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	return r.Queries.CountCompletedAnswersByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) CountTotalAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	return r.Queries.CountTotalAnswersByExercise(ctx, exerciseID)
}

// Attempts Repository methods
func (r *DatabaseRepository) CreateAttempt(ctx context.Context, arg codelabdb.CreateAttemptParams) error {
	return r.Queries.CreateAttempt(ctx, arg)
}

func (r *DatabaseRepository) GetAttempt(ctx context.Context, id int64) (codelabdb.Attempt, error) {
	return r.Queries.GetAttempt(ctx, id)
}

func (r *DatabaseRepository) GetAttemptsByAnswer(ctx context.Context, answerID int64) ([]codelabdb.Attempt, error) {
	return r.Queries.GetAttemptsByAnswer(ctx, answerID)
}

func (r *DatabaseRepository) GetAttemptsByUserAndExercise(ctx context.Context, arg codelabdb.GetAttemptsByUserAndExerciseParams) ([]codelabdb.Attempt, error) {
	return r.Queries.GetAttemptsByUserAndExercise(ctx, arg)
}

func (r *DatabaseRepository) GetLatestAttemptByAnswer(ctx context.Context, answerID int64) (codelabdb.Attempt, error) {
	return r.Queries.GetLatestAttemptByAnswer(ctx, answerID)
}

func (r *DatabaseRepository) CountAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error) {
	return r.Queries.CountAttemptsByAnswer(ctx, answerID)
}

func (r *DatabaseRepository) CountSuccessfulAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error) {
	return r.Queries.CountSuccessfulAttemptsByAnswer(ctx, answerID)
}

func (r *DatabaseRepository) CountTotalAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	return r.Queries.CountTotalAttemptsByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) CountSuccessfulAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	return r.Queries.CountSuccessfulAttemptsByExercise(ctx, exerciseID)
}

func (r *DatabaseRepository) GetAttemptsWithAnswerInfo(ctx context.Context, exerciseID int64) ([]codelabdb.GetAttemptsWithAnswerInfoRow, error) {
	return r.Queries.GetAttemptsWithAnswerInfo(ctx, exerciseID)
}

func (r *DatabaseRepository) GetUserAttemptsForExercise(ctx context.Context, arg codelabdb.GetUserAttemptsForExerciseParams) ([]codelabdb.Attempt, error) {
	return r.Queries.GetUserAttemptsForExercise(ctx, arg)
}
