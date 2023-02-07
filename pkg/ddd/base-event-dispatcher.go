package ddd

type IEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
