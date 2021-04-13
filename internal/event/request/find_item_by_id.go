package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type FindItemById struct {
	*framework.BaseEvent
	item int32
	itemId string
}

func (findItemById *FindItemById) ToMessage() interface{} {
	return nil
}

func (findItemById *FindItemById) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.FindItemById)
	findItemById.itemId = messagePb.GetItemId()
	findItemById.item = int32(messagePb.GetFindItem())
}