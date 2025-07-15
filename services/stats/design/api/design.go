package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("stats", func() {
	Title("Statistics Management API")
	Description("Microservice for managing statistics in sustainable classrooms")
	Version("1.0")

	// HTTP Server configuration
	Server("stats-http", func() {
		Description("Statistics microservice HTTP server")

		// Services available via HTTP
		Services("stats")

		// HTTP host configuration
		Host("localhost", func() {
			Description("Development HTTP server")
			URI("http://localhost:8080")
		})
	})

	// gRPC Server configuration
	Server("stats-grpc", func() {
		Description("Statistics microservice gRPC server")

		// Services available via gRPC
		Services("stats")

		// gRPC host configuration
		Host("localhost", func() {
			Description("Development gRPC server")
			URI("grpc://localhost:9090")
		})
	})

	// Global documentation
	Docs(func() {
		Description("Statistics API Documentation")
		URL("https://github.com/your-org/profiles-service")
	})
})