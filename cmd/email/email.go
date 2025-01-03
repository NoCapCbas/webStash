package email

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Email struct {
	To      string
	Subject string
	Body    string
}

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type MailgunServiceImpl struct{}

func (e *MailgunServiceImpl) SendEmail(to string, subject string, body string) error {
	log.Println("Sending email to", to, "with subject", subject, "and body", body)
	return nil
}

func main() {
	// redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//subscribe to the channel
	subscriber := redisClient.Subscribe(ctx, "email")
	channel := subscriber.Channel()

	//listen for messages
	for msg := range channel {
		log.Println(msg.Channel, msg.Payload)
		if msg.Channel == "send-email" {
			email := Email{}
			json.Unmarshal([]byte(msg.Payload), &email)
			SendEmail(email.To, email.Subject, email.Body)
		}
	}
}
