package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type AddResponse struct {
	*framework.BaseEvent
	ResponseStatus int32
	ObjectId string
	Error string
}

func (addResponse *AddResponse) ToMessage() interface{} {
	addResPb := &databaseGrpc.AddResponse{}
	addResPb.Error = addResponse.Error
	addResPb.ObjectId = addResponse.ObjectId
	addResPb.ResponseStatus = databaseGrpc.RESPONSE_STATUS(addResponse.ResponseStatus)
	return addResPb
}

func (addResponse *AddResponse) FromMessage(message interface{}) {}

func NewAddResponse(responseStatus int32, ObjectId string, err string) *BaseResponse {
	return &BaseResponse{
		AddResponse: &AddResponse{
			ResponseStatus: responseStatus,
			ObjectId:       ObjectId,
			Error:          err,
		},
	}
}