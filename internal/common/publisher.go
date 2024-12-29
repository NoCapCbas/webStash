package pubsub

import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "log"
)

// Publisher is responsible for sending events to Redis
type Publisher struct {
    redisClient *redis.Client
}

// NewPublisher initializes a new Publisher instance
func NewPublisher(redisClient *redis.Client) *Publisher {
    return &Publisher{
        redisClient: redisClient,
    }
}

// Publish sends an event to the specified channel
func (p *Publisher) Publish(ctx context.Context, channel string, event interface{}) error {
    // Marshal the event into JSON
    eventData, err := json.Marshal(event)
    if err != nil {
        return err
    }

    // Publish the event to the Redis channel
    err = p.redisClient.Publish(ctx, channel, eventData).Err()
    if err != nil {
        return err
    }

    log.Printf("Published event to channel %s: %s", channel, string(eventData))
    return nil
}

