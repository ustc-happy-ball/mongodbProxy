package event

type Event interface {
	ToMessage()
	FromMessage()
}
