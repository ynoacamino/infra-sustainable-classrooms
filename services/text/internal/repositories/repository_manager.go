package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
	"google.golang.org/grpc"
)

type RepositoryManager struct {
	CourseRepo          ports.CourseRepository
	ArticleRepo         ports.ArticleRepository
	SectionRepo         ports.SectionRepository
	ProfilesServiceRepo ports.ProfilesServiceRepository
	pool               *pgxpool.Pool
}

func NewRepositoryManager(
	pool *pgxpool.Pool, grpccoon *grpc.ClientConn,
) *RepositoryManager {
	return &RepositoryManager{
		CourseRepo:          NewCourseRepository(pool),
		ArticleRepo:         NewArticleRepository(pool),
		SectionRepo:         NewSectionRepository(pool),
		ProfilesServiceRepo: NewProfilesServiceRepository(grpccoon),
		pool :               pool,
	}
}

func (rm *RepositoryManager) Close() {
	rm.pool.Close()
}

func (rm *RepositoryManager) GetPool() *pgxpool.Pool {
	return rm.pool
}
