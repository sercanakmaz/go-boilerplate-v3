package ddd

import (
	"context"
	"sync"
)

var eventContextKey = "eventContext"

type EventContext struct {
	raisedEvents     []IBaseEvent
	dispatchedEvents []IBaseEvent
	mu               *sync.Mutex
}

func NewEventContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, eventContextKey, &EventContext{
		raisedEvents:     make([]IBaseEvent, 0),
		dispatchedEvents: make([]IBaseEvent, 0),
		mu:               &sync.Mutex{},
	})
}

func GetEventContext(ctx context.Context) *EventContext {
	var result = ctx.Value(eventContextKey)

	if result == nil {
		return nil
	}

	return result.(*EventContext)
}

func DispatchDomainEvents(ctx context.Context, order IAggregateRoot) {
	var eventContext = GetEventContext(ctx)

	if eventContext == nil {
		order.ClearDomainEvents()
		return
	}

	for _, event := range order.GetDomainEvents() {
		eventContext.AddRaised(event)
	}

	order.ClearDomainEvents()
}

func (self *EventContext) AddRaised(event IBaseEvent) {
	self.mu.Lock()
	self.raisedEvents = append(self.raisedEvents, event)
	self.mu.Unlock()
}

func (self *EventContext) AddDispatched(event IBaseEvent) {
	self.mu.Lock()
	self.dispatchedEvents = append(self.raisedEvents, event)
	self.mu.Unlock()
}

func (self *EventContext) TakeRaised() IBaseEvent {

	self.mu.Lock()

	if len(self.raisedEvents) == 0 {
		return nil
	}

	result := self.raisedEvents[0]

	self.raisedEvents = self.raisedEvents[1:]

	self.mu.Unlock()

	return result
}

func (self *EventContext) TakeDispatched() IBaseEvent {

	self.mu.Lock()

	if len(self.dispatchedEvents) == 0 {
		return nil
	}

	result := self.dispatchedEvents[0]

	self.dispatchedEvents = self.dispatchedEvents[1:]

	self.mu.Unlock()

	return result
}
