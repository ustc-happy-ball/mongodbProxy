package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type DeleteResponse struct {
	*framework.BaseEvent
	responseStatus int32
	error string
}

func (deleteResponse *DeleteResponse) ToMessage() interface{} {
	deleteResPb := &databaseGrpc.DeleteResponse{}
	deleteResPb.ResponseStatus = databaseGrpc.RESPONSE_STATUS(deleteResponse.responseStatus)
	deleteResPb.Error = deleteResponse.error
	return deleteResPb
}

func (deleteResponse *DeleteResponse) FromMessage(interface{}) {

}