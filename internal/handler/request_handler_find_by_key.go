package handler

import "github.com/TianqiS/database_for_happyball/framework/event"

type FindByKeyRequestHandler struct {}

var FindByKeyHandler = &FindByKeyRequestHandler{}

func (handler *FindByKeyRequestHandler) OnEvent(e event.Event) {

}
