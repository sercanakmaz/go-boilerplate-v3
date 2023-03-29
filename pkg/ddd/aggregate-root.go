package ddd

type IAggregateRoot interface {
	RaiseEvent(event IBaseEvent)
	GetDomainEvents() []IBaseEvent
	ClearDomainEvents()
}
