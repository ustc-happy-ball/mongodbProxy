package event

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/internal/event/response"
)

func NewAddResMsg(responseStatus int32, ObjectId string, err string) *DbMessage {
	addRes := response.NewAddResponse(responseStatus, ObjectId, err)
	newDbMsg := NewDbMessage(configs.ResponseAdd, addRes)
	return newDbMsg
}

func NewDeleteResMsg(responseStatus int32, err string) *DbMessage {
	deleteRes := response.NewDeleteResponse(responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseDelete, deleteRes)
	return newDbMsg
}

func NewFindResMsg(responseStatus int32, result interface{}, responseType int32, err string) *DbMessage {
	findRes := response.NewFindResponse(responseStatus, result, responseType, err)
	newDbMsg := NewDbMessage(configs.ResponseFind, findRes)
	return newDbMsg
}

func NewUpdateResMsg(responseStatus int32, err string) *DbMessage {
	updateRes := response.NewDeleteResponse(responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseUpdate, updateRes)
	return newDbMsg
}