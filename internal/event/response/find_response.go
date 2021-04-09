package response

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/internal/event/info"
	"github.com/TianqiS/database_for_happyball/model"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type FindResponse struct {
	*framework.BaseEvent
	responseStatus int32
	results interface{}
	responseType string
	error string
}

func (findResponse *FindResponse) ToMessage() interface{} {
	findResPb := &databaseGrpc.FindResponse{}
	findResPb.ResponseStatus = databaseGrpc.RESPONSE_STATUS(findResponse.responseStatus)
	resultAny, err := switcher(findResponse.results, findResponse.responseType)
	if err != nil {
		log.Println("find response error: ", err)
	}
	findResPb.Results = resultAny
	findResPb.Error = findResponse.error
	return findResPb
}

func (findResponse *FindResponse) FromMessage(interface{}) {

}

func switcher(findResult interface{}, resType string) (*anypb.Any, error){
	var result *anypb.Any
	var err error
	switch resType {
	case "account":
		resultModel := findResult.(*model.Account)
		resultEvent := &info.AccountEvent{}
		resultEvent.FromModel(resultModel)
		result, err = anypb.New(resultEvent.ToMessage().(*databaseGrpc.Account))
		break
	default:

	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
