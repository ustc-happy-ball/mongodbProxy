package framework

type BaseEvent struct {
	Code int32
}

func (be *BaseEvent) GetCode() int32 {
	return be.Code
}

func (be *BaseEvent) SetCode(code int32) {
	be.Code = code
}
