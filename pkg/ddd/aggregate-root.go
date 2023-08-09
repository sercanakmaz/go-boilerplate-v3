package ddd

import "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/event-handler"

type IAggregateRoot interface {
	RaiseEvent(event event_handler.IBaseEvent)
	GetDomainEvents() []event_handler.IBaseEvent
	ClearDomainEvents()
}
