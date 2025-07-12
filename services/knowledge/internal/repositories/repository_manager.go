package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
	"google.golang.org/grpc"
)

type RepositoryManager struct {
	TestRepo            ports.TestRepository
	QuestionRepo        ports.QuestionRepository
	SubmissionRepo      ports.SubmissionRepository
	TestCategoryRepo    ports.TestCategoryRepository
	AuthServiceRepo     ports.AuthServiceRepository
	ProfilesServiceRepo ports.ProfilesServiceRepository
	pool                *pgxpool.Pool
}

func NewRepositoryManager(pool *pgxpool.Pool, authGrpcConn *grpc.ClientConn, profilesGrpcConn *grpc.ClientConn) *RepositoryManager {
	return &RepositoryManager{
		TestRepo:            NewTestRepository(pool),
		QuestionRepo:        NewQuestionRepository(pool),
		SubmissionRepo:      NewSubmissionRepository(pool),
		TestCategoryRepo:    NewTestCategoryRepository(pool),
		AuthServiceRepo:     NewAuthServiceRepository(authGrpcConn),
		ProfilesServiceRepo: NewProfilesServiceRepository(profilesGrpcConn),
		pool:                pool,
	}
}

// Close cierra la conexión del pool
func (rm *RepositoryManager) Close() {
	rm.pool.Close()
}

// GetPool retorna el pool de conexiones para casos específicos
func (rm *RepositoryManager) GetPool() *pgxpool.Pool {
	return rm.pool
}
