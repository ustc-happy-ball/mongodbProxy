package event

type EventHandler interface {
	OnEvent(event Event)
}