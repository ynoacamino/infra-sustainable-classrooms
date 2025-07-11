package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"sync"

	profilespb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/grpc/profiles/pb"
	profilessvr "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/grpc/profiles/server"
	profiles "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handleGRPCServer(ctx context.Context, u *url.URL, profilesEndpoints *profiles.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var profilesServer *profilessvr.Server
	profilesServer = profilessvr.New(profilesEndpoints, nil)

	chain := grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx))
	if dbg {
		chain = grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx), debug.UnaryServerInterceptor())
	}

	srv := grpc.NewServer(chain)
	profilespb.RegisterProfilesServer(srv, profilesServer)

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
