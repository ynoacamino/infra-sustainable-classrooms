package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type SubmissionRepository interface {
	// Simplified submission operations
	CreateSubmission(ctx context.Context, params knowledgedb.CreateSubmissionParams) (knowledgedb.TestSubmission, error)
	GetSubmissionById(ctx context.Context, id int64) (knowledgedb.TestSubmission, error)
	GetUserSubmissions(ctx context.Context, userID int64) ([]knowledgedb.GetUserSubmissionsRow, error)
	CheckUserCompletedTest(ctx context.Context, params knowledgedb.CheckUserCompletedTestParams) (bool, error)
}
