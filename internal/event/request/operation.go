package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type Operation struct {
	framework.BaseEvent
	Op string
	Items []*MatchItem
}

func (operation *Operation) ToMessage() interface{} {
	return nil
}

func (operation *Operation) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.Operation)
	operation.Op = messagePb.Op
	itemsArr := make([]*MatchItem, len(messagePb.Items))
	for i, itemPb := range messagePb.Items {
		matchItem := &MatchItem{}
		matchItem.FromMessage(itemPb)
		itemsArr[i] = matchItem
	}
	operation.Items = itemsArr
}
