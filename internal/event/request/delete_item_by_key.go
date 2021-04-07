package request

import "github.com/TianqiS/database_for_happyball/framework"

type DeleteItemByKey struct {
	*framework.BaseEvent
}

func (deleteItemByKey *DeleteItemByKey) ToMessage() interface{} {
	return nil
}

func (deleteItemByKey *DeleteItemByKey) FromMessage(interface{}) {

}
