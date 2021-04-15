package event

import (
	"github.com/TianqiS/database_for_happyball/configs"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/internal/event/request"
	"github.com/TianqiS/database_for_happyball/internal/event/response"
)

type DbMessage struct {
	framework.BaseEvent
	MessageType int32
	MessageCode int32
	Request *request.BaseRequest
	Response *response.BaseResponse
	ResResult chan *DbMessage
	Error chan error
}

func (dbMessage *DbMessage) ToMessage() interface{} {
	dbMessagePb := &databaseGrpc.DbMessage{
		MessageCode: databaseGrpc.MESSAGE_CODE(dbMessage.MessageCode),
		MessageType: databaseGrpc.MESSAGE_TYPE(dbMessage.MessageType),
	}
	resEncoder(dbMessage, dbMessagePb)
	return dbMessagePb
}

func (dbMessage *DbMessage) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.DbMessage)
	req := switcher(messagePb)
	dbMessage.Request = req
	dbMessage.MessageType = int32(messagePb.MessageType)
	dbMessage.MessageCode = int32(messagePb.MessageCode)
	dbMessage.SetCode(dbMessage.MessageCode)
}

func resEncoder(dbMessage *DbMessage, dbMessagePb *databaseGrpc.DbMessage) {
	dbMessagePb.Response = &databaseGrpc.Response{}
	dbMessagePb.Response.ResponseStatus = databaseGrpc.RESPONSE_STATUS(dbMessage.Response.ResponseStatus)
	dbMessagePb.Response.Error = dbMessage.Response.Error
	if dbMessage.Response.ResponseStatus != configs.ResponseStatusOk {
		return
	}
	switch dbMessage.MessageCode {
	case configs.ResponseFind:
		dbMessagePb.Response.FindResponse = &databaseGrpc.FindResponse{}
		findRes := dbMessage.Response.FindResponse.ToMessage().(*databaseGrpc.FindResponse)
		dbMessagePb.Response.FindResponse = findRes
		break
	case configs.ResponseAdd:
		dbMessagePb.Response.AddResponse = &databaseGrpc.AddResponse{}
		resRes := dbMessage.Response.AddResponse.ToMessage().(*databaseGrpc.AddResponse)
		dbMessagePb.Response.AddResponse = resRes
		break
	case configs.ResponseUpdate:
		dbMessagePb.Response.UpdateResponse = &databaseGrpc.UpdateResponse{}
		updateRes := dbMessage.Response.UpdateResponse.ToMessage().(*databaseGrpc.UpdateResponse)
		dbMessagePb.Response.UpdateResponse = updateRes
		break
	case configs.ResponseDelete:
		dbMessagePb.Response.DeleteResponse = &databaseGrpc.DeleteResponse{}
		deleteRes := dbMessage.Response.DeleteResponse.ToMessage().(*databaseGrpc.DeleteResponse)
		dbMessagePb.Response.DeleteResponse = deleteRes
		break
	}
}

func switcher(messagePb *databaseGrpc.DbMessage) *request.BaseRequest {
	req := &request.BaseRequest{}
	switch int32(messagePb.GetMessageCode()) {
	case configs.RequestFindById:
		reqPb := messagePb.GetRequest().GetFindItemById()
		req.FindItemById = &request.FindItemById{}
		req.FindItemById.FromMessage(reqPb)
		break
	case configs.RequestFindByKey:
		reqPb := messagePb.GetRequest().GetFindItemByKey()
		req.FindItemByKey = &request.FindItemByKey{}
		req.FindItemByKey.FromMessage(reqPb)
		break
	case configs.RequestAddRequest:
		reqPb := messagePb.GetRequest().GetAddItem()
		req.AddItem = &request.AddItem{}
		req.AddItem.FromMessage(reqPb)
		break
	case configs.RequestDeleteById:
		reqPb := messagePb.GetRequest().GetDeleteItemById()
		req.DeleteItemById = &request.DeleteItemById{}
		req.DeleteItemById.FromMessage(reqPb)
		break
	case configs.RequestDeleteByKey:
		reqPb := messagePb.GetRequest().GetDeleteItemByKey()
		req.DeleteItemByKey = &request.DeleteItemByKey{}
		req.DeleteItemByKey.FromMessage(reqPb)
		break
	case configs.RequestUpdateById:
		reqPb := messagePb.GetRequest().GetUpdateItemById()
		req.UpdateItemById = &request.UpdateItemById{}
		req.UpdateItemById.FromMessage(reqPb)
		break
	case configs.RequestUpdateByKey:
		reqPb := messagePb.GetRequest().GetUpdateItemByKey()
		req.UpdateItemByKey = &request.UpdateItemByKey{}
		req.UpdateItemByKey.FromMessage(reqPb)
		break
	default:
	}
	return req
}

func NewDbMessage(messageCode int32, res *response.BaseResponse) *DbMessage {
	return &DbMessage{
		MessageType: configs.MessageTypeResponse,
		MessageCode: messageCode,
		Response:    res,
	}
}



