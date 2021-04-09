package request

import (
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateItemById struct {
	*framework.BaseEvent
	objectId string
	updateItem int32
	items map[string]interface{}
}

func (updateItemById *UpdateItemById) ToMessage() interface{} {
	return nil
}

func (updateItemById *UpdateItemById) FromMessage(interface{}) {

}
