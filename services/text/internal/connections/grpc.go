package connections

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectGRPCConfig interface {
	GrpcAddress() string
}

func ConnectGRPC(cfg ConnectGRPCConfig) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return conn, nil
}
