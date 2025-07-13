package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("mailing", func() {
	Description("Mailing microservice for sending emails")

	HTTP(func() {
		Path("/mailing")
	})

	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	// Global error definitions for the service
	Error("invalid_input", String, "Invalid input parameters")
	Error("email_send_failed", String, "Failed to send email")
	Error("smtp_connection_failed", String, "Failed to connect to SMTP server")
	Error("service_unavailable", String, "Service temporarily unavailable")

	// Send email method
	Method("SendEmail", func() {
		Description("Send an email message via SMTP")

		Payload(func() {
			Field(1, "email", EmailMessage, "Email message to send")
			Required("email")
		})

		Result(EmailResponse)

		HTTP(func() {
			POST("/send")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("email_send_failed", StatusInternalServerError)
			Response("smtp_connection_failed", StatusServiceUnavailable)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("invalid_input", CodeInvalidArgument)
			Response("email_send_failed", CodeInternal)
			Response("smtp_connection_failed", CodeUnavailable)
			Response("service_unavailable", CodeUnavailable)
		})
	})
})
