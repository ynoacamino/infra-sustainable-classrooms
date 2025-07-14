package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the stats microservice API
var _ = API("stats", func() {
	Title("Statistics Management API")
	Description("Microservice for managing user progress statistics in sustainable classrooms")
	Version("1.0")

	// HTTP Server configuration
	Server("stats-http", func() {
		Description("Stats microservice HTTP server")

		// Services available via HTTP
		Services("stats")

		// HTTP host configuration
		Host("localhost", func() {
			Description("Development HTTP server")
			URI("http://localhost:8081")
		})
	})

	// gRPC Server configuration
	Server("stats-grpc", func() {
		Description("Stats microservice gRPC server")

		// Services available via gRPC
		Services("stats")

		// gRPC host configuration
		Host("localhost", func() {
			Description("Development gRPC server")
			URI("grpc://localhost:9091")
		})
	})

	// Global documentation
	Docs(func() {
		Description("Stats API Documentation")
		URL("https://github.com/your-org/stats-service")
	})
})
