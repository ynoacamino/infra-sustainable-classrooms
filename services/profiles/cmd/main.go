package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/config"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/connections"
	profilesapi "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/controllers"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/repositories"
	"goa.design/clue/debug"
	"goa.design/clue/log"
)

func main() {
	fmt.Println("---------------------------------------------------------")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(context.Background(), fmt.Errorf("failed to load config: %w", err))
	}

	ctx := cfg.Ctx
	log.Printf(ctx, "starting auth service in %s mode", cfg.Environment)
	log.Printf(ctx, "configured for reverse proxy deployment - serving plain HTTP")

	pool, err := connections.ConnectDB(cfg)
	if err != nil {
		log.Fatal(ctx, fmt.Errorf("failed to connect to database: %w", err))
	}
	defer pool.Close()
	log.Printf(ctx, "database connection established")

	grpccoon, err := connections.ConnectGRPC(config.ConnectGRPCConfig{
		GrpcAddress: cfg.AuthGRPCAddress,
	})

	// Initialize repository manager
	reposManager := repositories.NewRepositoryManager(pool, grpccoon)
	defer reposManager.Close()

	// Initialize service with repository manager
	var profilesSvc profiles.Service = profilesapi.NewProfiles(reposManager)

	var profilesEndpoints *profiles.Endpoints
	profilesEndpoints = profiles.NewEndpoints(profilesSvc)
	profilesEndpoints.Use(debug.LogPayloads())
	profilesEndpoints.Use(log.Endpoint)

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	{
		addr := fmt.Sprintf("http://0.0.0.0:%s", cfg.HTTPPort)
		u, err := url.Parse(addr)
		if err != nil {
			log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
		}
		handleHTTPServer(ctx, u, profilesEndpoints, &wg, errc, cfg.Debug)
	}

	{
		addr := fmt.Sprintf("grpc://0.0.0.0:%s", cfg.GRPCPort)
		u, err := url.Parse(addr)
		if err != nil {
			log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
		}
		handleGRPCServer(ctx, u, profilesEndpoints, &wg, errc, cfg.Debug)
	}

	log.Printf(ctx, "exiting (%v)", <-errc)
	cancel()
	wg.Wait()
	log.Printf(ctx, "exited")
}
