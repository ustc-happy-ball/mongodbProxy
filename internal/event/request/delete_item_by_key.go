package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteItemByKey struct {
	framework.BaseEvent
	DeleteItem int32
	MatchItem []*MatchItem
}

func (deleteItemByKey *DeleteItemByKey) ToMessage() interface{} {
	return nil
}

func (deleteItemByKey *DeleteItemByKey) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DeleteItemByKey)
	deleteItemByKey.DeleteItem = int32(messagePb.GetDeleteItem())
	matchItemArr := messagePb.GetMatchItem()
	deleteItemByKey.MatchItem = make([]*MatchItem, len(matchItemArr))
	for i, itemPb := range matchItemArr {
		item := &MatchItem{}
		item.FromMessage(itemPb)
		deleteItemByKey.MatchItem[i] = item
	}
}
