package connections

import (
	"fmt"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectGRPC(cfg config.ConnectGRPCConfig) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return conn, nil
}
