package ports

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type TestCategoryRepository interface {
	// Category CRUD operations
	CreateCategory(ctx context.Context, params knowledgedb.CreateCategoryParams) error
	GetCategoryById(ctx context.Context, id int32) (knowledgedb.TestCategory, error)
	GetAllCategories(ctx context.Context) ([]knowledgedb.TestCategory, error)
	UpdateCategory(ctx context.Context, params knowledgedb.UpdateCategoryParams) (knowledgedb.TestCategory, error)
	DeleteCategory(ctx context.Context, id int32) error

	// Category statistics
	CountTestsByCategory(context.Context, pgtype.Int4) (int64, error)
	GetCategoriesWithTestCounts(ctx context.Context) ([]knowledgedb.GetCategoriesWithTestCountsRow, error)
}
