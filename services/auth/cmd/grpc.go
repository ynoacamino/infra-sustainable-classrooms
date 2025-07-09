package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"sync"

	auth "github.com/ynoacamino/infrastructure/services/auth/gen/auth"
	authpb "github.com/ynoacamino/infrastructure/services/auth/gen/grpc/auth/pb"
	authsvr "github.com/ynoacamino/infrastructure/services/auth/gen/grpc/auth/server"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handleGRPCServer(ctx context.Context, u *url.URL, authEndpoints *auth.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var authServer *authsvr.Server
	authServer = authsvr.New(authEndpoints, nil)

	chain := grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx))
	if dbg {
		chain = grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx), debug.UnaryServerInterceptor())
	}

	srv := grpc.NewServer(chain)
	authpb.RegisterAuthServer(srv, authServer)

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
