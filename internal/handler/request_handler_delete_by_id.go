package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type DeleteByIdRequestHandler struct {}

var DeleteByIdHandler = &DeleteByIdRequestHandler{}

func (handler *DeleteByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}