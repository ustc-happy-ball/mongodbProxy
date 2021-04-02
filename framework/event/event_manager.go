package event

import "sync"

type EventManager struct {
	rw         *sync.RWMutex
	handlerMap map[int64]EventHandler
	eventMap   map[int64]Event
}

var EManager = newManager()

func newManager() *EventManager {
	return &EventManager{
		rw:         &sync.RWMutex{},
		handlerMap: make(map[int64]EventHandler),
		eventMap:   make(map[int64]Event),
	}
}

func (this *EventManager) Register(msgCode int64, event Event, handler EventHandler) {
	if nil != this.handlerMap[msgCode] {
		return
	}
	if nil != handler {
		this.rw.Lock()
		this.handlerMap[msgCode] = handler
		this.rw.Unlock()
	}
	if nil != this.eventMap[msgCode] {
		return
	}
	if nil != event {
		this.rw.Lock()
		this.eventMap[msgCode] = event
		this.rw.Unlock()
	}
}

func (this *EventManager) FetchHandler(msgCode int64) EventHandler {
	return this.handlerMap[msgCode]
}

func (this *EventManager) FetchEvent(msgCode int64) Event {
	return this.eventMap[msgCode]
}

