package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type AddResponse struct {
	*framework.BaseEvent
	ObjectId string
}

func (addResponse *AddResponse) ToMessage() interface{} {
	addResPb := &databaseGrpc.AddResponse{}
	addResPb.ObjectId = addResponse.ObjectId
	return addResPb
}

func (addResponse *AddResponse) FromMessage(message interface{}) {}

func NewAddResponse(ObjectId string, resStatus int32, err string) *BaseResponse {
	return &BaseResponse{
		ResponseStatus:  resStatus,
		Error: err,
		AddResponse: &AddResponse{
			ObjectId:       ObjectId,
		},
	}
}