package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type TestRepository interface {
	// Simplified Test CRUD operations
	CreateTest(ctx context.Context, params knowledgedb.CreateTestParams) error
	GetTestById(ctx context.Context, id int64) (knowledgedb.Test, error)
	GetMyTests(ctx context.Context, createdBy int64) ([]knowledgedb.Test, error)
	GetAvailableTests(ctx context.Context, createdBy int64) ([]knowledgedb.GetAvailableTestsRow, error)
	UpdateTest(ctx context.Context, params knowledgedb.UpdateTestParams) error
	DeleteTest(ctx context.Context, id int64) error
}
