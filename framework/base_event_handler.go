package framework

import "github.com/TianqiS/database_for_happyball/framework/event"

type BaseEventHandler struct {

}

var EVENT_HANDLER = &BaseEventHandler{}

func (be *BaseEventHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	handler := event.EManager.FetchHandler(e.GetCode())
	if nil != handler {
		handler.OnEvent(e)
	}
}