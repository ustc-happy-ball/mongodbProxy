package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type DeleteByKeyRequestHandler struct {}

var DeleteByKeyHandler = &DeleteByKeyRequestHandler{}

func (handler *DeleteByKeyRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}
