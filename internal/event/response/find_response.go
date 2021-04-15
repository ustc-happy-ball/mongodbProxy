package response

import (
	"github.com/TianqiS/database_for_happyball/configs"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/internal/event/info"
	"github.com/TianqiS/database_for_happyball/model"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type FindResponse struct {
	framework.BaseEvent
	Results interface{}
	item int32
}

func (findResponse *FindResponse) ToMessage() interface{} {
	findResPb := &databaseGrpc.FindResponse{}
	resultAny, err := switcher(findResponse.Results, findResponse.item)
	if err != nil {
		log.Println("find response error: ", err)
	}
	findResPb.Results = resultAny
	return findResPb
}

func (findResponse *FindResponse) FromMessage(interface{}) {

}

func switcher(findResult interface{}, resType int32) (*anypb.Any, error){
	var result *anypb.Any
	var err error
	switch resType {
	case configs.ItemPlayer:
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

func NewFindResponse(result interface{}, itemType int32, responseStatus int32, err string) *BaseResponse {
	return &BaseResponse{
		ResponseStatus: responseStatus,
		Error: err,
		FindResponse: &FindResponse{
			Results:        result,
			item:   itemType,
		},
	}
}
