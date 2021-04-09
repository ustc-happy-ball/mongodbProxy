package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type UpdateResponse struct {
	*framework.BaseEvent
	responseStatus int32
	error string
}

func (updateResponse *UpdateResponse) ToMessage() interface{} {
	return &databaseGrpc.UpdateResponse{
		ResponseStatus: databaseGrpc.RESPONSE_STATUS(updateResponse.responseStatus),
		Error: updateResponse.error,
	}
}

func (updateResponse *UpdateResponse) FromMessage(interface{}) {

}