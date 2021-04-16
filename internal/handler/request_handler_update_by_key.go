package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type UpdateByKeyRequestHandler struct {}

var UpdateByKeyHandler = &UpdateByKeyRequestHandler{}

func (handler *UpdateByKeyRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	updateByKeyRequest := dbMessageEvent.Request.UpdateItemByKey
	itemType := updateByKeyRequest.UpdateItem
	collection := db.CollectionMap[itemType]
	err := collection.UpdateItemByKey(updateByKeyRequest.MatchItem, updateByKeyRequest.Operation)
	if err != nil {
		dbMessageEvent.Error <- err
		return
	}
	updateResult := event2.NewUpdateResMsg(configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- updateResult
}
