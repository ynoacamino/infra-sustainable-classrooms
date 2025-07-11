package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the profiles microservice API
var _ = API("profiles", func() {
	Title("User Profiles Management API")
	Description("Microservice for managing student and teacher profiles in sustainable classrooms")
	Version("1.0")

	// HTTP Server configuration
	Server("profiles-http", func() {
		Description("Profiles microservice HTTP server")

		// Services available via HTTP
		Services("profiles")

		// HTTP host configuration
		Host("localhost", func() {
			Description("Development HTTP server")
			URI("http://localhost:8082")
		})
	})

	// gRPC Server configuration
	Server("profiles-grpc", func() {
		Description("Profiles microservice gRPC server")

		// Services available via gRPC
		Services("profiles")

		// gRPC host configuration
		Host("localhost", func() {
			Description("Development gRPC server")
			URI("grpc://localhost:9092")
		})
	})

	// Global documentation
	Docs(func() {
		Description("User Profiles Management API Documentation")
		URL("https://github.com/your-org/profiles-service")
	})
})
