package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateResponse struct {
	*framework.BaseEvent
	ResponseStatus int32
	Error string
}

func (updateResponse *UpdateResponse) ToMessage() interface{} {
	return &databaseGrpc.UpdateResponse{
		ResponseStatus: databaseGrpc.RESPONSE_STATUS(updateResponse.ResponseStatus),
		Error: updateResponse.Error,
	}
}

func (updateResponse *UpdateResponse) FromMessage(interface{}) {

}

func NewUpdateResponse(responseStatus int32, err string) *BaseResponse {
	return &BaseResponse{
		UpdateResponse: &UpdateResponse{
			ResponseStatus: responseStatus,
			Error:          err,
		},
	}
}