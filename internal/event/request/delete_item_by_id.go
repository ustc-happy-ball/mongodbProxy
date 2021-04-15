package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteItemById struct {
	framework.BaseEvent
	DeleteItem int32
	ItemId string
}

func (deleteItemById *DeleteItemById) ToMessage() interface{} {
	return nil
}

func (deleteItemById *DeleteItemById) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DeleteItemById)
	deleteItemById.DeleteItem = int32(messagePb.GetDeleteItem())
	deleteItemById.ItemId = messagePb.GetItemId()
}


