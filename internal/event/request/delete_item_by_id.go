package request

import "github.com/TianqiS/database_for_happyball/framework"

type DeleteItemById struct {
	*framework.BaseEvent
}

func (deleteItemById *DeleteItemById) ToMessage() interface{} {
	return nil
}

func (deleteItemById *DeleteItemById) FromMessage(interface{}) {

}


