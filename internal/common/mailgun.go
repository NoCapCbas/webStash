package common

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient struct {
	mg     mailgun.Mailgun
	sender string
	domain string
}

func NewMailgunClient() *MailgunClient {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	sender := os.Getenv("MAILGUN_SENDER") // e.g., "WebStash <noreply@yourdomain.com>"

	mg := mailgun.NewMailgun(domain, apiKey)

	return &MailgunClient{
		mg:     mg,
		sender: sender,
		domain: domain,
	}
}

func (m *MailgunClient) SendMagicLink(recipient, magicLink string) error {
	// Create template data
	data := struct {
		MagicLink string
	}{
		MagicLink: magicLink,
	}

	// Parse the email template
	tmpl, err := template.ParseFiles("templates/mail/magic-link.html")
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	// Execute the template with the data
	var htmlBody bytes.Buffer
	if err := tmpl.ExecuteTemplate(&htmlBody, "magic-link-email", data); err != nil {
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	// Create the message
	message := mailgun.NewMessage(
		m.sender,
		"Sign in to webStash",
		"", // Plain text version (optional)
		recipient,
	)
	message.SetHTML(htmlBody.String())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the email
	_, id, err := m.mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	// Log the message ID for tracking
	fmt.Printf("Email sent successfully. Message ID: %s\n", id)
	return nil
}
