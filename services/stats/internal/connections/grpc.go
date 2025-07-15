package connections

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCConnections holds all gRPC client connections
type GRPCConnections struct {
	TextServiceConn     *grpc.ClientConn
	ProfilesServiceConn *grpc.ClientConn
}

// NewGRPCConnections creates new gRPC connections to other services
func NewGRPCConnections() (*GRPCConnections, error) {
	// Get service endpoints from environment variables
	textServiceAddr := getEnvOrDefault("TEXT_SERVICE_GRPC_ADDR", "text:8080")
	profilesServiceAddr := getEnvOrDefault("PROFILES_SERVICE_GRPC_ADDR", "profiles:8080")

	// Determine if we should use TLS
	useTLS := getEnvOrDefault("USE_TLS", "false") == "true"

	var creds credentials.TransportCredentials
	if useTLS {
		// Use TLS credentials
		creds = credentials.NewTLS(&tls.Config{
			ServerName: "text", // Adjust according to your certificate
		})
	} else {
		// Use insecure credentials for development
		creds = insecure.NewCredentials()
	}

	// Create connection to text service
	textConn, err := grpc.Dial(textServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to text service at %s: %w", textServiceAddr, err)
	}

	// Create connection to profiles service
	profilesConn, err := grpc.Dial(profilesServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		textConn.Close() // Clean up the text connection
		return nil, fmt.Errorf("failed to connect to profiles service at %s: %w", profilesServiceAddr, err)
	}

	log.Printf("Successfully connected to text service at %s", textServiceAddr)
	log.Printf("Successfully connected to profiles service at %s", profilesServiceAddr)

	return &GRPCConnections{
		TextServiceConn:     textConn,
		ProfilesServiceConn: profilesConn,
	}, nil
}

// Close closes all gRPC connections
func (g *GRPCConnections) Close() error {
	var errs []error

	if g.TextServiceConn != nil {
		if err := g.TextServiceConn.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close text service connection: %w", err))
		}
	}

	if g.ProfilesServiceConn != nil {
		if err := g.ProfilesServiceConn.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close profiles service connection: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}

	return nil
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}