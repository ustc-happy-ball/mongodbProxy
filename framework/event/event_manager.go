package event

import "sync"

type EventManager struct {
	rw         *sync.RWMutex
	handlerMap map[int32]EventHandler
}

var EManager = newManager()

func newManager() *EventManager {
	return &EventManager{
		rw:         &sync.RWMutex{},
		handlerMap: make(map[int32]EventHandler),
	}
}

func (this *EventManager) Register(msgCode int32, handler EventHandler) {
	if nil != this.handlerMap[msgCode] {
		return
	}
	if nil != handler {
		this.rw.Lock()
		this.handlerMap[msgCode] = handler
		this.rw.Unlock()
	}
}

func (this *EventManager) FetchHandler(msgCode int32) EventHandler {
	return this.handlerMap[msgCode]
}


