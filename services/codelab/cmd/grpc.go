package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"sync"

	codelab "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/codelab"
	codelabpb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/grpc/codelab/pb"
	codelabsvr "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/grpc/codelab/server"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handleGRPCServer(ctx context.Context, u *url.URL, codelabEndpoints *codelab.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var codelabServer *codelabsvr.Server
	codelabServer = codelabsvr.New(codelabEndpoints)

	chain := grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx))
	if dbg {
		chain = grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx), debug.UnaryServerInterceptor())
	}

	srv := grpc.NewServer(chain)
	codelabpb.RegisterCodelabServer(srv, codelabServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			log.Printf(ctx, "serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			if lis == nil {
				errc <- fmt.Errorf("failed to listen on %q", u.Host)
			}
			log.Printf(ctx, "gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
