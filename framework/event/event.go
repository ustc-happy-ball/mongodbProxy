package event

type Event interface {
	GetCode() int32
	SetCode(int32)
	ToMessage() interface{}
	FromMessage(interface{})
}
