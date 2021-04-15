package handler

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	event2 "github.com/TianqiS/database_for_happyball/internal/event"
)

type FindByKeyRequestHandler struct {}

var FindByKeyHandler = &FindByKeyRequestHandler{}

func (handler *FindByKeyRequestHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	dbMessageEvent := e.(*event2.DbMessage)
	findByKeyRequest := dbMessageEvent.Request.FindItemByKey
	itemType := dbMessageEvent.Request.FindItemByKey.FindItem
	collection := db.CollectionMap[itemType]
	res, err := collection.FindItemsByKey(findByKeyRequest.MatchItem)
	if err != nil {
		dbMessageEvent.Error <- err
		return
	}
	findResponseMsg := event2.NewFindResMsg(res, itemType, configs.ResponseStatusOk, "")
	dbMessageEvent.ResResult <- findResponseMsg
}
