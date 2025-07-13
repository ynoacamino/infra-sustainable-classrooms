package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/config"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/connections"
	videolearningapi "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/controllers"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/repositories"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/services"
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
	log.Printf(ctx, "starting video_learning service in %s mode", cfg.Environment)
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
	if err != nil {
		log.Fatal(ctx, fmt.Errorf("failed to connect to gRPC server: %w", err))
	}
	defer grpccoon.Close()

	minioClient, err := connections.ConnectMinio(cfg)
	if err != nil {
		log.Fatal(ctx, fmt.Errorf("failed to connect to MinIO: %w", err))
	}
	log.Printf(ctx, "MinIO connection established")

	redisClient, err := connections.ConnectRedis(cfg)
	if err != nil {
		log.Fatal(ctx, fmt.Errorf("failed to connect to Redis: %w", err))
	}
	defer redisClient.Close()
	log.Printf(ctx, "Redis connection established")

	// Initialize repository manager
	reposManager := repositories.NewRepositoryManager(pool, grpccoon, redisClient, minioClient)
	defer reposManager.Close()
	// Initialize aggregation service for periodic cache-to-database sync
	aggregationService := services.NewAggregationService(cfg.Ctx, reposManager)

	// Start aggregation service with configured interval
	aggregationInterval := time.Duration(cfg.AggregationIntervalSeconds) * time.Second
	log.Printf(ctx, "Starting aggregation service with interval: %v", aggregationInterval)
	aggregationService.Start(aggregationInterval)
	defer aggregationService.Stop()

	// Initialize service with repository manager
	var videoLearningSvc videolearning.Service = videolearningapi.NewVideoLearning(reposManager)

	var videoLearningEndpoints = videolearning.NewEndpoints(videoLearningSvc)
	videoLearningEndpoints.Use(debug.LogPayloads())
	videoLearningEndpoints.Use(log.Endpoint)

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
		handleHTTPServer(ctx, u, videoLearningEndpoints, &wg, errc, cfg.Debug)
	}

	{
		addr := fmt.Sprintf("grpc://0.0.0.0:%s", cfg.GRPCPort)
		u, err := url.Parse(addr)
		if err != nil {
			log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
		}
		handleGRPCServer(ctx, u, videoLearningEndpoints, &wg, errc, cfg.Debug)
	}

	log.Printf(ctx, "exiting (%v)", <-errc)
	cancel()
	wg.Wait()
	log.Printf(ctx, "exited")
}
