package repositories

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/ports"
	"google.golang.org/grpc"
)

type RepositoryManager struct {
	TextServiceRepo     ports.TextServiceRepository
	ProfilesServiceRepo ports.ProfilesServiceRepository
}

func NewRepositoryManager(
	textGrpcConn *grpc.ClientConn, 
	profilesGrpcConn *grpc.ClientConn,
) *RepositoryManager {
	return &RepositoryManager{
		TextServiceRepo:     NewTextServiceRepository(textGrpcConn),
		ProfilesServiceRepo: NewProfilesServiceRepository(profilesGrpcConn),
	}
}

func (rm *RepositoryManager) Close() {
	// Close connections if needed
}
