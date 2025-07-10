package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/ports"
)

type RepositoryManager struct {
	ProfileRepo        ports.ProfileRepository
	StudentProfileRepo ports.StudentProfileRepository
	TeacherProfileRepo ports.TeacherProfileRepository
	pool               *pgxpool.Pool
}

func NewRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		ProfileRepo:        NewProfileRepository(pool),
		StudentProfileRepo: NewStudentProfileRepository(pool),
		TeacherProfileRepo: NewTeacherProfileRepository(pool),
		pool:               pool,
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
