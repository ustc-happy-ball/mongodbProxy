package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type UpdateByIdRequestHandler struct {}

var UpdateByIdHandler = &UpdateByIdRequestHandler{}

func (handler *UpdateByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	updateByIdRequest := dbMessageEvent.Request.UpdateItemById
	itemType := updateByIdRequest.UpdateItem
	collection := db.CollectionMap[itemType]
	err := collection.UpdateItemById(updateByIdRequest.ObjectId, updateByIdRequest.Operation)
	if err != nil {
		dbMessageEvent.Error <- err
	}
	updateResponse := event2.NewUpdateResMsg(configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- updateResponse
}
