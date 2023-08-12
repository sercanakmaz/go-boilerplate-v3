package ddd

import "context"

type IBaseEventHandler[E IBaseEvent] interface {
	Handle(ctx context.Context, event E) error
}
