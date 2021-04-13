package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteItemById struct {
	deleteItem int32
	itemId string
	*framework.BaseEvent
}

func (deleteItemById *DeleteItemById) ToMessage() interface{} {
	return nil
}

func (deleteItemById *DeleteItemById) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DeleteItemById)
	deleteItemById.deleteItem = int32(messagePb.GetDeleteItem())
	deleteItemById.itemId = messagePb.GetItemId()
}


