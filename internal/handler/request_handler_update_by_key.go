package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type UpdateByKeyRequestHandler struct {}

var UpdateByKeyHandler = &UpdateByKeyRequestHandler{}

func (handler *UpdateByKeyRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}
