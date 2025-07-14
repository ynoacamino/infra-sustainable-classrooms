package design

import (
	. "goa.design/goa/v3/dsl"
)

// EmailMessage represents an email message to be sent
var EmailMessage = Type("EmailMessage", func() {
	Description("Email message structure")

	Field(1, "to", ArrayOf(String), "Recipient email addresses", func() {
		MinLength(1)
		Example([]string{"user@example.com", "admin@example.com"})
	})

	Field(2, "cc", ArrayOf(String), "Carbon copy email addresses", func() {
		Example([]string{"cc@example.com"})
	})

	Field(3, "bcc", ArrayOf(String), "Blind carbon copy email addresses", func() {
		Example([]string{"bcc@example.com"})
	})

	Field(4, "subject", String, "Email subject", func() {
		MinLength(1)
		MaxLength(200)
		Example("Welcome to our platform!")
	})

	Field(5, "body", String, "Email body content", func() {
		MinLength(1)
		Example("Hello! Welcome to our platform. We're excited to have you on board.")
	})

	Field(6, "is_html", Boolean, "Whether the body content is HTML", func() {
		Default(false)
		Example(true)
	})

	Required("to", "subject", "body")
})

// EmailResponse represents the response after sending an email
var EmailResponse = Type("EmailResponse", func() {
	Description("Response after sending an email")

	Field(1, "success", Boolean, "Whether the email was sent successfully", func() {
		Example(true)
	})

	Field(2, "message", String, "Response message", func() {
		Example("Email sent successfully")
	})

	Field(3, "message_id", String, "Message ID from the SMTP server", func() {
		Example("abc123@mail.example.com")
	})

	Required("success", "message")
})
