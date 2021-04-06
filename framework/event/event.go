package event

type Event interface {
	GetCode() int64
	SetCode(int64)
	ToMessage() interface{}
	FromMessage(interface{})
}
