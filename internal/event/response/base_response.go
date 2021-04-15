package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type BaseResponse struct {
	*framework.BaseEvent
	ResponseStatus int32
	Error string
	*AddResponse
	*DeleteResponse
	*FindResponse
	*UpdateResponse
}

func (baseResponse *BaseResponse) ToMessage() interface{} {
	baseResponsePb := &databaseGrpc.Response{
		ResponseStatus: databaseGrpc.RESPONSE_STATUS(baseResponse.ResponseStatus),
		Error: baseResponse.Error,
	}
	if baseResponse.AddResponse != nil {
		baseResponsePb.AddResponse = baseResponse.AddResponse.ToMessage().(*databaseGrpc.AddResponse)
		return baseResponsePb
	}

	if baseResponse.DeleteResponse != nil {
		baseResponsePb.DeleteResponse = baseResponse.DeleteResponse.ToMessage().(*databaseGrpc.DeleteResponse)
		return baseResponsePb
	}

	if baseResponse.FindResponse != nil {
		baseResponsePb.FindResponse = baseResponse.FindResponse.ToMessage().(*databaseGrpc.FindResponse)
		return baseResponsePb
	}

	if baseResponse.UpdateResponse != nil {
		baseResponsePb.UpdateResponse = baseResponse.UpdateResponse.ToMessage().(*databaseGrpc.UpdateResponse)
		return baseResponsePb
	}
	return baseResponsePb
}

func (baseResponse *BaseResponse) FromMessage(message interface{}) {}

func NewErrResponse(responseStatus int32, err string) *BaseResponse {
	return &BaseResponse{ResponseStatus: responseStatus, Error: err}
}
