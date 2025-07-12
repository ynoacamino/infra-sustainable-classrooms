package repositories

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
	"google.golang.org/grpc"

	genauth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	genAuthclient "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/grpc/auth/client"
)

type AuthServiceRepository struct {
}

func NewAuthServiceRepository(conn *grpc.ClientConn) ports.AuthServiceRepository {
	grpcClient := genAuthclient.NewClient(conn)

	client := &genauth.Client{
		ValidateUserEndpoint: grpcClient.ValidateUser(),
		GetUserByIDEndpoint:  grpcClient.GetUserByID(),
	}

	return client
}
