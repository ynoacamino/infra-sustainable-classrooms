package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type SubmissionRepository interface {
	// Test submission CRUD operations
	CreateSubmission(ctx context.Context, params knowledgedb.CreateSubmissionParams) (knowledgedb.TestSubmission, error)
	GetSubmissionById(ctx context.Context, id int64) (knowledgedb.TestSubmission, error)
	GetSubmissionByUserAndTest(ctx context.Context, params knowledgedb.GetSubmissionByUserAndTestParams) (knowledgedb.TestSubmission, error)
	CompleteSubmission(ctx context.Context, params knowledgedb.CompleteSubmissionParams) (knowledgedb.TestSubmission, error)

	// User submissions
	GetUserSubmissions(ctx context.Context, params knowledgedb.GetUserSubmissionsParams) ([]knowledgedb.GetUserSubmissionsRow, error)

	// Test submissions by test
	GetTestParticipants(ctx context.Context, params knowledgedb.GetTestParticipantsParams) ([]knowledgedb.GetTestParticipantsRow, error)
	CountTestParticipants(ctx context.Context, testID int64) (int64, error)

	// Submission results
	GetSubmissionWithAnswers(ctx context.Context, id int64) (knowledgedb.GetSubmissionWithAnswersRow, error)

	// Submission validation
	CheckExistingSubmission(ctx context.Context, params knowledgedb.CheckExistingSubmissionParams) (int64, error)
	CheckUserCompletedTest(ctx context.Context, params knowledgedb.CheckUserCompletedTestParams) (bool, error)
}
