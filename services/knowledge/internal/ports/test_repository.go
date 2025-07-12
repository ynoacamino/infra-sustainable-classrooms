package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type TestRepository interface {
	CheckTestAccess(ctx context.Context, id int64) (knowledgedb.CheckTestAccessRow, error)
	CountTestsByCreator(ctx context.Context, createdBy int64) (int64, error)
	CreateTest(ctx context.Context, arg knowledgedb.CreateTestParams) error
	DeleteTest(ctx context.Context, arg knowledgedb.DeleteTestParams) error
	GetAvailableTests(ctx context.Context, arg knowledgedb.GetAvailableTestsParams) ([]knowledgedb.GetAvailableTestsRow, error)
	GetExpiredTests(ctx context.Context) ([]knowledgedb.Test, error)
	GetPopularTests(ctx context.Context, arg knowledgedb.GetPopularTestsParams) ([]knowledgedb.GetPopularTestsRow, error)
	GetTestById(ctx context.Context, id int64) (knowledgedb.GetTestByIdRow, error)
	GetTestPreview(ctx context.Context, id int64) (knowledgedb.GetTestPreviewRow, error)
	GetTestWithQuestions(ctx context.Context, id int64) (knowledgedb.GetTestWithQuestionsRow, error)
	GetTestsByCategory(ctx context.Context, arg knowledgedb.GetTestsByCategoryParams) ([]knowledgedb.GetTestsByCategoryRow, error)
	GetTestsByCreator(ctx context.Context, arg knowledgedb.GetTestsByCreatorParams) ([]knowledgedb.GetTestsByCreatorRow, error)
	GetTestsWithSubmissionCounts(ctx context.Context, arg knowledgedb.GetTestsWithSubmissionCountsParams) ([]knowledgedb.GetTestsWithSubmissionCountsRow, error)
	UpdateTest(ctx context.Context, arg knowledgedb.UpdateTestParams) error
	UpdateTestActiveStatus(ctx context.Context, arg knowledgedb.UpdateTestActiveStatusParams) error
}
