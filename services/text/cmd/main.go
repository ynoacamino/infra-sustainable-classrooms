package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/config"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/connections"
	textapi "github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/controllers"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/repositories"
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
	log.Printf(ctx, "starting text service in %s mode", cfg.Environment)
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
	
	reposManager := repositories.NewRepositoryManager(pool, grpccoon)

	// Initialize repository manager
	repoManager := repositories.NewRepositoryManager(pool, grpcConn)
	defer repoManager.Close()

	// Initialize service with repository manager
	var textSvc text.Service = textapi.NewText(repoManager)

	var textEndpoints *text.Endpoints
	textEndpoints = text.NewEndpoints(textSvc)
	textEndpoints.Use(debug.LogPayloads())
	textEndpoints.Use(log.Endpoint)

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
}
