package event_handler

type IEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
