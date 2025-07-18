// Code generated by goa v3.21.1, DO NOT EDIT.
//
// codelab gRPC client
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/design/api
// -o ./services/codelab/

package client

import (
	codelabpb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/grpc/codelab/pb"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli codelabpb.CodelabClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the codelab service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: codelabpb.NewCodelabClient(cc),
		opts:    opts,
	}
}
