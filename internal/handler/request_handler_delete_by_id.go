package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type DeleteByIdRequestHandler struct {}

var DeleteByIdHandler = &DeleteByIdRequestHandler{}

func (handler *DeleteByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	deleteByIdRequest := dbMessageEvent.Request.DeleteItemById
	itemType := deleteByIdRequest.DeleteItem
	collection := db.CollectionMap[itemType]
	err := collection.DeleteItemById(deleteByIdRequest.ItemId)
	if err != nil {
		dbMessageEvent.Error <- err
	}
	deleteResponse := event2.NewDeleteResMsg(configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- deleteResponse
}