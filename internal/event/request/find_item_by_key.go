package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type FindItemByKey struct {
	framework.BaseEvent
	FindItem int32
	MatchItem []*MatchItem
}

func (findItemByKey *FindItemByKey) ToMessage() interface{} {
	return nil
}

func (findItemByKey *FindItemByKey) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.FindItemByKey)
	findItemByKey.FindItem = int32(messagePb.FindItem)
	matchItemArr := messagePb.GetMatchItem()
	findItemByKey.MatchItem = make([]*MatchItem, len(matchItemArr))
	for i, itemPb := range matchItemArr {
		item := &MatchItem{}
		item.FromMessage(itemPb)
		findItemByKey.MatchItem[i] = item
	}
}
