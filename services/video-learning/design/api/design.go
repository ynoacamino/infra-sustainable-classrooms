package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the global properties of the Video Learning API
var _ = API("video-learning", func() {
	Title("Video Learning Microservice")

	Description("Microservice for video learning platform with HTTP and gRPC support")

	Version("1.0")

	// CORS policy for frontend access
	HTTP(func() {
		// CORS configuration will be handled at the server level
	})

	// Server configuration
	Server("video-learning", func() {
		Description("Video Learning service server")

		// HTTP server
		Services("video-learning")
		Host("localhost", func() {
			Description("Development server")

			URI("http://localhost:8080")
		})
	})

	// gRPC server
	Server("grpc", func() {
		Description("Video Learning service gRPC server")

		Services("video-learning")

		Host("localhost", func() {
			Description("gRPC server for inter-service communication")
			URI("grpc://localhost:8090")
		})
	})
})
