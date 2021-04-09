package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
)

type AddResponse struct {
	*framework.BaseEvent
	responseStatus int32
	objectId string
	error string
}

func (addResponse *AddResponse) ToMessage() interface{} {
	addResPb := &databaseGrpc.AddResponse{}
	addResPb.Error = addResponse.error
	addResPb.ObjectId = addResponse.objectId
	addResPb.ResponseStatus = databaseGrpc.RESPONSE_STATUS(addResponse.responseStatus)
	return addResPb
}

func (addResponse *AddResponse) FromMessage(message interface{}) {}