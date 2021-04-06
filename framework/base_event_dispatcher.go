package framework

import "github.com/TianqiS/database_for_happyball/framework/event"

type BaseEventDispatcher struct {}

func (eDispatcher *BaseEventDispatcher) FireEvent(e event.Event) {
	EVENT_HANDLER.OnEvent(e)
}

func(eDispatcher *BaseEventDispatcher) Close() {

}