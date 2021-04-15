package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type FindItemById struct {
	*framework.BaseEvent
	Item int32
	ItemId string
}

func (findItemById *FindItemById) ToMessage() interface{} {
	return nil
}

func (findItemById *FindItemById) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.FindItemById)
	findItemById.ItemId = messagePb.GetItemId()
	findItemById.Item = int32(messagePb.GetFindItem())
}