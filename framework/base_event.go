package framework

type BaseEvent struct {
	Code int64
}

func (be *BaseEvent) GetCode() int64 {
	return be.Code
}

func (be *BaseEvent) SetCode(code int64) {
	be.Code = code
}
