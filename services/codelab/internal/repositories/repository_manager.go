package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/internal/ports"
	"google.golang.org/grpc"
)

type RepositoryManager struct {
	ProfilesServiceRepo ports.ProfilesServiceRepository
	AnswersRepo         ports.AnswersRepository
	ExercisesRepo       ports.ExercisesRepository
	TestsRepo           ports.TestsRepository
	AttemptsRepo        ports.AttemptsRepository
	pool                *pgxpool.Pool
}

func NewRepositoryManager(
	pool *pgxpool.Pool, grpccoon *grpc.ClientConn,
) *RepositoryManager {
	return &RepositoryManager{
		ProfilesServiceRepo: NewProfilesServiceRepository(grpccoon),
		AnswersRepo:         NewAnswersRepository(pool),
		ExercisesRepo:       NewExercisesRepository(pool),
		TestsRepo:           NewTestsRepository(pool),
		AttemptsRepo:        NewAttemptsRepository(pool),
		pool:                pool,
	}
}

func (rm *RepositoryManager) Close() {
	rm.pool.Close()
}

func (rm *RepositoryManager) GetPool() *pgxpool.Pool {
	return rm.pool
}
