package events

import "time"

type Message interface {
	Type() string
}

// Will be sent by nat
type CreatedFeedMessage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// The contrat with the interface

func (m CreatedFeedMessage) Type() string {
	return "created_feed"
}
