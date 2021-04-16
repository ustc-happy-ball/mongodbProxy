package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateItemById struct {
	framework.BaseEvent
	ObjectId string
	UpdateItem int32
	Operation *Operation
}

func (updateItemById *UpdateItemById) ToMessage() interface{} {
	return nil
}

func (updateItemById *UpdateItemById) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.UpdateItemById)
	updateItemById.ObjectId = messagePb.ObjectId
	updateItemById.UpdateItem = int32(messagePb.UpdateItem)
	op := &Operation{}
	op.FromMessage(messagePb.Operation)
	updateItemById.Operation = op
}
