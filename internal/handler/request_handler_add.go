package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type addRequestHandler struct {}

var AddHandler = &addRequestHandler{}

func (handler *addRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	addItemRequest := dbMessageEvent.Request.AddItem
	itemType := addItemRequest.AddItem
	collection := db.CollectionMap[itemType]
	objectId, err := collection.InsertItem(addItemRequest.Item)
	if err != nil {
		dbMessageEvent.Error <- err
		return
	}
	addItemResponseMsg := event2.NewAddResMsg(objectId, configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- addItemResponseMsg
}
