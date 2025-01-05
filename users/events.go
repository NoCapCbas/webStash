package users

import (
	"encoding/json"
	"log"
	"time"
)

// EventHandler processes incoming events and triggers appropriate actions
type EventHandler struct {
	userService UserService
}

func NewEventHandler(userService UserService) *EventHandler {
	return &EventHandler{
		userService: userService,
	}
}

type EventType string

const (
	UserCreatedEvent  EventType = "user.created"
	UserVerifiedEvent EventType = "user.verified"
	UserUpdatedEvent  EventType = "user.updated"
)

// HandleMessage processes incoming messages and triggers appropriate events
func (h *EventHandler) HandleMessage(message string) error {
	var event struct {
		Type      EventType `json:"event_type"`
		Timestamp time.Time `json:"timestamp"`
		// Source    string      `json:"source"`
		Payload interface{} `json:"payload"`
	}

	if err := json.Unmarshal([]byte(message), &event); err != nil {
		return err
	}

	switch event.Type {
	case UserCreatedEvent:
		return h.handleSignup(event.Payload)
	case UserVerifiedEvent:
		return h.handleVerify(event.Payload)
	case UserUpdatedEvent:
		return h.handleUpdate(event.Payload)
	default:
		log.Printf("Unknown event type: %s", event.Type)
		return nil
	}
}

func (h *EventHandler) handleSignup(payload interface{}) error {
	// Convert payload to appropriate type
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	// Process signup
	_, err = h.userService.SignUp(&user)
	return err
}

func (h *EventHandler) handleVerify(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	return h.userService.Verify(&user)
}

func (h *EventHandler) handleUpdate(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	return h.userService.Update(&user)
}
