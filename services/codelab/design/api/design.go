package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the global properties of the Codelab API
var _ = API("codelab", func() {
	Title("Codelab Microservice")

	Description("Microservice for coding exercises, tests, answers and attempts with HTTP and gRPC support")

	Version("1.0")

	// CORS policy for frontend access
	HTTP(func() {
		// CORS configuration will be handled at the server level
	})

	// Server configuration
	Server("codelab", func() {
		Description("Codelab service server")

		// HTTP server
		Services("codelab")
		Host("localhost", func() {
			Description("Development server")

			URI("http://localhost:8080")
		})
	})

	// gRPC server
	Server("grpc", func() {
		Description("Codelab service gRPC server")

		Services("codelab")

		Host("localhost", func() {
			Description("gRPC server for inter-service communication")
			URI("grpc://localhost:8090")
		})
	})
})
