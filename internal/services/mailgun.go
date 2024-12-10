package services

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// MailService defines the interface for sending emails
type MailService interface {
	SendMagicLink(recipient, magicLink string) error
}

// MailgunService implements the MailService interface
type MailgunService struct {
	mg     mailgun.Mailgun
	sender string
	domain string
}

// NewMailgunService creates a new instance of MailgunService
func NewMailgunService() MailService {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	sender := os.Getenv("MAILGUN_SENDER") // e.g., "WebStash <noreply@yourdomain.com>"

	mg := mailgun.NewMailgun(domain, apiKey)

	return &MailgunService{
		mg:     mg,
		sender: sender,
		domain: domain,
	}
}

// SendMagicLink sends a magic link email to the recipient
func (m *MailgunService) SendMagicLink(recipient, magicLink string) error {
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
