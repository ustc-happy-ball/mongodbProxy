package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type addRequestHandler struct {}

var AddHandler = &addRequestHandler{}

func (handler *addRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
}
