package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the profiles microservice API
var _ = API("text", func() {
	Title("Text Management API")
	Description("Microservice for managing text based content in sustainable classrooms")
	Version("1.0")

	// HTTP Server configuration
	Server("text-http", func() {
		Description("Text microservice HTTP server")

		// Services available via HTTP
		Services("text")

		// HTTP host configuration
		Host("localhost", func() {
			Description("Development HTTP server")
			URI("http://localhost:8080")
		})
	})

	// gRPC Server configuration
	Server("text-grpc", func() {
		Description("Text microservice gRPC server")

		// Services available via gRPC
		Services("text")

		// gRPC host configuration
		Host("localhost", func() {
			Description("Development gRPC server")
			URI("grpc://localhost:9090")
		})
	})

	// Global documentation
	Docs(func() {
		Description("Text API Documentation")
		URL("https://github.com/your-org/profiles-service")
	})
})
