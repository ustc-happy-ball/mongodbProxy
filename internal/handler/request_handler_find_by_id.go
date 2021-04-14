package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type FindByIdRequestHandler struct {}

var FindByIdHandler = &FindByIdRequestHandler{}

func (handler *FindByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}
