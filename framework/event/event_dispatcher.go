package event

type EventDispatcher interface {
	FireEvent()
	Close()
}
