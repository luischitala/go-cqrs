package events

import (
	"context"

	"github.com/luischitala/go/cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	// Callback
	OnCreateFeed(f func(CreatedFeedMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}

func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)
}

func SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}

func OnCreateFeed(ctx context.Context, f func(CreatedFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}
