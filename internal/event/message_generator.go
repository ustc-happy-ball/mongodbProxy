package event

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/internal/event/response"
)

func NewAddResMsg(ObjectId string, responseStatus int32,  err string) *DbMessage {
	addRes := response.NewAddResponse(ObjectId, responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseAdd, addRes)
	return newDbMsg
}

func NewDeleteResMsg(responseStatus int32, err string) *DbMessage {
	deleteRes := response.NewDeleteResponse(responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseDelete, deleteRes)
	return newDbMsg
}

func NewFindResMsg(result interface{}, itemType int32, responseStatus int32, err string) *DbMessage {
	findRes := response.NewFindResponse(result, itemType, responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseFind, findRes)
	return newDbMsg
}

func NewUpdateResMsg(responseStatus int32, err string) *DbMessage {
	updateRes := response.NewDeleteResponse(responseStatus, err)
	newDbMsg := NewDbMessage(configs.ResponseUpdate, updateRes)
	return newDbMsg
}

func NewErrResMsg(responseCode int32, responseStatus int32, err string) *DbMessage {
	errRes := response.NewErrResponse(responseStatus, err)
	newDbMsg := NewDbMessage(responseCode, errRes)
	return newDbMsg
}