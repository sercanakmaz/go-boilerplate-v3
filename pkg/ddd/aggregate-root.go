package ddd

type IAggregateRoot interface {
	AddEvent(event IBaseEvent)
	GetDomainEvents() []IBaseEvent
	ClearDomainEvents()
}
