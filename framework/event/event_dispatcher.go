package event

type EventDispatcher interface {
	FireEvent(event Event)
	Close()
}
