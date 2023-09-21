package ddd

import (
	"context"
)

type IEventPublisher interface {
	Publish(ctx context.Context) error
}
