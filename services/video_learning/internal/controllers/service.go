package controllers

import (
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/repositories"
)

type videolearningsrvc struct {
	userCategoryLikeRepo ports.UserCategoryLikeRepository
	videoCategoryRepo    ports.VideoCategoryRepository
	videoCommentRepo     ports.VideoCommentRepository
	videoRepo            ports.VideoRepository
	videoTagRepo         ports.VideoTagRepository
	profileServiceRepo   ports.ProfilesServiceRepository
	cacheRepo            ports.CacheRepository
	storageRepo          ports.StorageRepository
	repoManager          *repositories.RepositoryManager
}

func NewVideoLearning(
	repoManager *repositories.RepositoryManager,
) videolearning.Service {
	return &videolearningsrvc{
		userCategoryLikeRepo: repoManager.UserCategoryLikeRepo,
		videoCategoryRepo:    repoManager.VideoCategoryRepo,
		videoCommentRepo:     repoManager.VideoCommentRepo,
		videoRepo:            repoManager.VideoRepo,
		videoTagRepo:         repoManager.VideoTagRepo,
		profileServiceRepo:   repoManager.ProfilesServiceRepo,
		cacheRepo:            repoManager.CacheRepo,
		storageRepo:          repoManager.StorageRepo,
		repoManager:          repoManager,
	}
}
