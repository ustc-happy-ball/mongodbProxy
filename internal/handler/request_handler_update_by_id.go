package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type UpdateByIdRequestHandler struct {}

var UpdateByIdHandler = &UpdateByIdRequestHandler{}

func (handler *UpdateByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}
