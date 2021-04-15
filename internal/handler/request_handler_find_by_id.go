package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type FindByIdRequestHandler struct {}

var FindByIdHandler = &FindByIdRequestHandler{}

func (handler *FindByIdRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	findByIdRequest := dbMessageEvent.Request.FindItemById
	itemType := dbMessageEvent.Request.FindItemById.Item
	collection := db.CollectionMap[findByIdRequest.Item]
	res, err := collection.FindOneItemById(findByIdRequest.ItemId)
	if err != nil {
		dbMessageEvent.Error <- err
		return
	}
	findResult := collection.GetModel()
	err = res.Decode(findResult)
	if err != nil {
		dbMessageEvent.Error <- err
		return
	}
	findResponseMsg := event2.NewFindResMsg([]interface{}{findResult}, itemType, configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- findResponseMsg
}
