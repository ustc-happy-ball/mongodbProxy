package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type DeleteByKeyRequestHandler struct {}

var DeleteByKeyHandler = &DeleteByKeyRequestHandler{}

func (handler *DeleteByKeyRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	deleteByKeyRequest := dbMessageEvent.Request.DeleteItemByKey
	itemType := deleteByKeyRequest.DeleteItem
	collection := db.CollectionMap[itemType]
	err := collection.DeleteItemByKey(deleteByKeyRequest.MatchItem)
	if err != nil {
		dbMessageEvent.Error <- err
	}
	deleteResponseMsg := event2.NewDeleteResMsg(configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- deleteResponseMsg
}
