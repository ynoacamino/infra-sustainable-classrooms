package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the global properties of the Auth API
var _ = API("auth", func() {
	Title("Authentication Microservice")

	Description("Microservice for authentication using OTP strategy with HTTP and gRPC support")

	Version("1.0")

	// CORS policy for frontend access
	HTTP(func() {
		// CORS configuration will be handled at the server level
	})

	// Server configuration
	Server("auth", func() {
		Description("Auth service server")

		// HTTP server
		Services("auth")
		Host("localhost", func() {
			Description("Development server")

			URI("http://localhost:8080")
		})
	})

	// gRPC server
	Server("grpc", func() {
		Description("Auth service gRPC server")

		Services("auth")

		Host("localhost", func() {
			Description("gRPC server for inter-service communication")
			URI("grpc://localhost:8090")
		})
	})
})
