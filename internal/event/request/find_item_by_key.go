package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type FindItemByKey struct {
	framework.BaseEvent
	deleteItem int32
	matchItem []*MatchItem
}

func (findItemByKey *FindItemByKey) ToMessage() interface{} {
	return nil
}

func (findItemByKey *FindItemByKey) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DeleteItemByKey)
	findItemByKey.deleteItem = int32(messagePb.DeleteItem)
	matchItemArr := messagePb.GetMatchItem()
	findItemByKey.matchItem = make([]*MatchItem, len(matchItemArr))
	for i, itemPb := range matchItemArr {
		item := &MatchItem{}
		item.FromMessage(itemPb)
		findItemByKey.matchItem[i] = item
	}
}
