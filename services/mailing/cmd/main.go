package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/config"
	mailing "github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/mailing"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/connections"
	mailingapi "github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/controllers"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/repositories"
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
	log.Printf(ctx, "starting mailing service in %s mode", cfg.Environment)
	log.Printf(ctx, "configured for reverse proxy deployment - serving plain HTTP")

	// Connect to SMTP server
	smtpClient, err := connections.ConnectSMTP(cfg)
	if err != nil {
		log.Fatal(ctx, fmt.Errorf("failed to connect to SMTP server: %w", err))
	}
	log.Printf(ctx, "SMTP connection established")

	// Initialize repository manager
	reposManager := repositories.NewRepositoryManager(smtpClient, cfg.SMTPFrom, cfg.IsProduction())
	defer reposManager.Close()

	// Initialize service with repository manager
	var mailingSvc mailing.Service = mailingapi.NewMailing(reposManager)

	var mailingEndpoints = mailing.NewEndpoints(mailingSvc)
	mailingEndpoints.Use(debug.LogPayloads())
	mailingEndpoints.Use(log.Endpoint)

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
		handleHTTPServer(ctx, u, mailingEndpoints, &wg, errc, cfg.Debug)
	}

	{
		addr := fmt.Sprintf("grpc://0.0.0.0:%s", cfg.GRPCPort)
		u, err := url.Parse(addr)
		if err != nil {
			log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
		}
		handleGRPCServer(ctx, u, mailingEndpoints, &wg, errc, cfg.Debug)
	}

	log.Printf(ctx, "exiting (%v)", <-errc)
	cancel()
	wg.Wait()
	log.Printf(ctx, "exited")
}
