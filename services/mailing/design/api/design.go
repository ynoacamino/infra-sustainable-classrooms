package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the global properties of the Mailing API
var _ = API("mailing", func() {
	Title("Mailing Microservice")

	Description("Microservice for sending emails via SMTP with gRPC support")

	Version("1.0")

	// Server configuration
	Server("mailing", func() {
		Description("Mailing service server")

		Services("mailing")
		Host("localhost", func() {
			Description("Development server")
			URI("http://localhost:8080")
		})
	})

	// gRPC server
	Server("grpc", func() {
		Description("Mailing service gRPC server")

		Services("mailing")

		Host("localhost", func() {
			Description("gRPC server for inter-service communication")
			URI("grpc://localhost:8090")
		})
	})
})
