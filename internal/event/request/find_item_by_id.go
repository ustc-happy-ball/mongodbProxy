package request

import "github.com/TianqiS/database_for_happyball/framework"

type FindItemById struct {
	*framework.BaseEvent
	item int32
	itemId string
}

func (findItemById *FindItemById) ToMessage() interface{} {
	return nil
}

func (findItemById *FindItemById) FromMessage(interface{}) {

}