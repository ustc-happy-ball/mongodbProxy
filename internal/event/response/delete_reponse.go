package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteResponse struct {
	*framework.BaseEvent
	ResponseStatus int32
	Error string
}

func (deleteResponse *DeleteResponse) ToMessage() interface{} {
	deleteResPb := &databaseGrpc.DeleteResponse{}
	deleteResPb.ResponseStatus = databaseGrpc.RESPONSE_STATUS(deleteResponse.ResponseStatus)
	deleteResPb.Error = deleteResponse.Error
	return deleteResPb
}

func (deleteResponse *DeleteResponse) FromMessage(interface{}) {

}