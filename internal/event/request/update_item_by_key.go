package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateItemByKey struct {
	framework.BaseEvent
	UpdateItem int32 // 更新的item的种类
	MatchItem []*MatchItem
	Operation *Operation
}

func (updateItemByKey *UpdateItemByKey) ToMessage() interface{} {
	return nil
}

func (updateItemByKey *UpdateItemByKey) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.UpdateItemByKey)
	updateItemByKey.UpdateItem = int32(messagePb.UpdateItem)
	matchItemArr := messagePb.GetMatchItem()
	updateItemByKey.MatchItem = make([]*MatchItem, len(matchItemArr))
	for i, itemPb := range matchItemArr {
		item := &MatchItem{}
		item.FromMessage(itemPb)
		updateItemByKey.MatchItem[i] = item
	}
	op := &Operation{}
	op.FromMessage(messagePb.Operation)
	updateItemByKey.Operation = op
}
