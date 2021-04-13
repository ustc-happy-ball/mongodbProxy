package request

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type MatchItem struct {
	*framework.BaseEvent
	key string
	matchVal interface{}
}

func (findItemByKey *MatchItem) ToMessage() interface{} {
	return nil
}

func (findItemByKey *MatchItem) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.MatchItem)
	findItemByKey.key = messagePb.GetKey()
	findItemByKey.matchVal =parseOneOf(messagePb)
}

func parseOneOf(matchItem *databaseGrpc.MatchItem) interface{} {
	matchItemMatch := matchItem.GetMatch()
	switch matchItemMatch.(type) {
	case *databaseGrpc.MatchItem_ValString:
		return matchItemMatch.(*databaseGrpc.MatchItem_ValString).ValString
	case *databaseGrpc.MatchItem_ValInt:
		return matchItemMatch.(*databaseGrpc.MatchItem_ValInt).ValInt
	case *databaseGrpc.MatchItem_ValBool:
		return matchItemMatch.(*databaseGrpc.MatchItem_ValBool).ValBool
	case *databaseGrpc.MatchItem_ValFloat:
		return matchItemMatch.(*databaseGrpc.MatchItem_ValFloat).ValFloat
	case *databaseGrpc.MatchItem_ValDouble:
		return matchItemMatch.(*databaseGrpc.MatchItem_ValDouble).ValDouble
	default:
		return nil
	}
}