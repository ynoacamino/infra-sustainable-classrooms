package repositories

import (
	"google.golang.org/grpc"

	genProfilesclient "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/grpc/profiles/client"
	genprofiles "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type ProfilesServiceRepository struct {
}

func NewProfilesServiceRepository(grpcConn *grpc.ClientConn) ports.ProfilesServiceRepository {
    grpcClient := genProfilesclient.NewClient(grpcConn)

    client := &genprofiles.Client{
        GetCompleteProfileEndpoint:   grpcClient.GetCompleteProfile(),
        GetPublicProfileByIDEndpoint: grpcClient.GetPublicProfileByID(),
        ValidateUserRoleEndpoint:     grpcClient.ValidateUserRole(),
    }

    return client
}