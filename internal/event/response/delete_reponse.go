package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteResponse struct {
	*framework.BaseEvent
}

func (deleteResponse *DeleteResponse) ToMessage() interface{} {
	deleteResPb := &databaseGrpc.DeleteResponse{}
	return deleteResPb
}

func (deleteResponse *DeleteResponse) FromMessage(interface{}) {

}

func NewDeleteResponse(responseStatus int32, err string) *BaseResponse {
	return &BaseResponse{
		ResponseStatus: responseStatus,
		Error: err,
		DeleteResponse: &DeleteResponse{
		},
	}
}