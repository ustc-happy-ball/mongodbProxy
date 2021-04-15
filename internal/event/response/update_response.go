package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateResponse struct {
	*framework.BaseEvent
}

func (updateResponse *UpdateResponse) ToMessage() interface{} {
	return &databaseGrpc.UpdateResponse{
	}
}

func (updateResponse *UpdateResponse) FromMessage(interface{}) {

}

func NewUpdateResponse(responseStatus int32, err string) *BaseResponse {
	return &BaseResponse{
		UpdateResponse: &UpdateResponse{
		},
	}
}