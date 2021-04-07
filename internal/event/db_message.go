package event

import (
	"github.com/TianqiS/database_for_happyball/framework"
)

type DbMessage struct {
	*framework.BaseEvent
	MessageType int32
	MessageCode int32
}

func (dbMessage *DbMessage) ToMessage() interface{} {
	return nil
}

func (dbMessage *DbMessage) FromMessage(interface{}) {

}



