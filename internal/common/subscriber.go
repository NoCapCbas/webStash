package common

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type EventHandler interface {
	HandleMessage(message string) error
}

type Message struct {
	Channel string
	Payload string
}

// String implements the Stringer interface for Message
func (m *Message) String() string {
	return fmt.Sprintf("Channel: %s, Payload: %s", m.Channel, m.Payload)
}

type Subscriber struct {
	client *redis.Client
	pubsub *redis.PubSub
}

func NewSubscriber(client *redis.Client) *Subscriber {
	return &Subscriber{
		client: client,
	}
}

// Subscribe subscribes to the specified channels and returns a channel for messages
func (s *Subscriber) Subscribe(ctx context.Context, channels ...string) error {
	s.pubsub = s.client.Subscribe(ctx, channels...)
	_, err := s.pubsub.Receive(ctx)
	return err
}

// GetMessage receives a message from the subscribed channels
func (s *Subscriber) GetMessage(ctx context.Context) (*Message, error) {
	msg, err := s.pubsub.ReceiveMessage(ctx)
	if err != nil {
		return nil, err
	}

	message := &Message{
		Channel: msg.Channel,
		Payload: msg.Payload,
	}

	log.Printf("Received message: %s", message)
	return message, nil
}

// Start a goroutine to listen for messages
func (s *Subscriber) Start(ctx context.Context, handler EventHandler) {
	go func() {
		for {
			message, err := s.GetMessage(ctx)
			if err != nil {
				log.Printf("Failed to get message: %v", err)
			}

			handler.HandleMessage(message.Payload)
		}
	}()
}

// Close closes the subscription
func (s *Subscriber) Close() error {
	return s.pubsub.Close()
}
