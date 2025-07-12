package design

import (
	. "goa.design/goa/v3/dsl"
)

// API defines the knowledge microservice API
var _ = API("knowledge", func() {
	Title("Knowledge Test Management API")

	Description("Microservice for managing MCQ tests, validations, grading, and student progress tracking")

	Version("1.0")

	Server("knowledge-http", func() {
		Description("Knowledge microservice HTTP server")

		Services("knowledge")

		Host("localhost", func() {
			Description("Development HTTP server")
			URI("http://localhost:8080")
		})
	})

	Server("knowledge-grpc", func() {
		Description("Knowledge microservice gRPC server")

		Services("knowledge")

		Host("localhost", func() {
			Description("Development gRPC server")
			URI("grpc://localhost:9090")
		})
	})

})
