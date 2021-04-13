package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteItemByKey struct {
	deleteItem int32
	matchItem []*MatchItem
	*framework.BaseEvent
}

func (deleteItemByKey *DeleteItemByKey) ToMessage() interface{} {
	return nil
}

func (deleteItemByKey *DeleteItemByKey) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DeleteItemByKey)
	deleteItemByKey.deleteItem = int32(messagePb.GetDeleteItem())
	matchItemArr := messagePb.GetMatchItem()
	deleteItemByKey.matchItem = make([]*MatchItem, len(matchItemArr))
	for i, itemPb := range matchItemArr {
		item := &MatchItem{}
		item.FromMessage(itemPb)
		deleteItemByKey.matchItem[i] = item
	}
}
