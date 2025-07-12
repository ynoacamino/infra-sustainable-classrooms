package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// RepositoryManager coordina todos los repositorios del dominio video_learning
type RepositoryManager struct {
	VideoRepo            ports.VideoRepository
	VideoCategoryRepo    ports.VideoCategoryRepository
	VideoCommentRepo     ports.VideoCommentRepository
	VideoTagRepo         ports.VideoTagRepository
	UserCategoryLikeRepo ports.UserCategoryLikeRepository
	CacheRepo            ports.CacheRepository
	StorageRepo          ports.StorageRepository
	TxManager            TransactionManager
	pool                 *pgxpool.Pool
}

// NewRepositoryManager crea una nueva instancia del manejador de repositorios
func NewRepositoryManager(pool *pgxpool.Pool, redisClient *redis.Client, minioClient *minio.Client) *RepositoryManager {
	return &RepositoryManager{
		VideoRepo:            NewVideoRepository(pool),
		VideoCategoryRepo:    NewVideoCategoryRepository(pool),
		VideoCommentRepo:     NewVideoCommentRepository(pool),
		VideoTagRepo:         NewVideoTagRepository(pool),
		UserCategoryLikeRepo: NewUserCategoryLikeRepository(pool),
		CacheRepo:            NewCacheRepository(redisClient),
		StorageRepo:          NewStorageRepository(minioClient),
		TxManager:            NewTransactionManager(pool),
		pool:                 pool,
	}
}

// Close cierra la conexión del pool
func (rm *RepositoryManager) Close() {
	rm.pool.Close()
}

// WithTransaction crea repositorios transaccionales
func (rm *RepositoryManager) WithTransaction(tx pgx.Tx) *TransactionalRepositories {
	return &TransactionalRepositories{
		VideoRepo:            rm.NewVideoRepositoryWithTx(tx),
		VideoCategoryRepo:    rm.NewVideoCategoryRepositoryWithTx(tx),
		VideoCommentRepo:     rm.NewVideoCommentRepositoryWithTx(tx),
		VideoTagRepo:         rm.NewVideoTagRepositoryWithTx(tx),
		UserCategoryLikeRepo: rm.NewUserCategoryLikeRepositoryWithTx(tx),
	}
}

// TransactionalRepositories agrupa repositorios transaccionales
type TransactionalRepositories struct {
	VideoRepo            ports.VideoRepository
	VideoCategoryRepo    ports.VideoCategoryRepository
	VideoCommentRepo     ports.VideoCommentRepository
	VideoTagRepo         ports.VideoTagRepository
	UserCategoryLikeRepo ports.UserCategoryLikeRepository
}

// Factory methods para repositorios transaccionales
func (rm *RepositoryManager) NewVideoRepositoryWithTx(tx pgx.Tx) ports.VideoRepository {
	return NewVideoRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewVideoCategoryRepositoryWithTx(tx pgx.Tx) ports.VideoCategoryRepository {
	return NewVideoCategoryRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewVideoCommentRepositoryWithTx(tx pgx.Tx) ports.VideoCommentRepository {
	return NewVideoCommentRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewVideoTagRepositoryWithTx(tx pgx.Tx) ports.VideoTagRepository {
	return NewVideoTagRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewUserCategoryLikeRepositoryWithTx(tx pgx.Tx) ports.UserCategoryLikeRepository {
	return NewUserCategoryLikeRepositoryWithTx(tx)
}
