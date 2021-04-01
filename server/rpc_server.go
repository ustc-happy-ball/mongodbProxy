package server

import (
	"context"
	"fmt"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type RpcServer struct {
	databaseGrpc.UnimplementedDbServiceServer
}

func GetServer() *RpcServer {
	return &RpcServer{}
}

func (s *RpcServer) SendRequest(ctx context.Context, in *databaseGrpc.DbMessage) (*databaseGrpc.DbMessage, error) {
	if in.GetRequest().GetFindItemById() != nil {
		acc, err := db.AccountCollection.FindAccount(in.GetRequest().GetFindItemById().ItemId)
		if err != nil {
			log.Fatal("查找时发生了错误", err)
		}
		accPb := databaseGrpc.Account{
			Name:          acc.Name,
			LoginPassword: acc.LoginPassword,
			Level:         acc.Level,
			Delete:        acc.Delete,
			Region:        acc.Region,
			Phone:         acc.Phone,
			CreateAt:      0,
			UpdateAt:      0,
		}
		accAny, _ := anypb.New(&accPb)
		fmt.Println("获取的acc为", acc)
		return &databaseGrpc.DbMessage{
			MessageType: databaseGrpc.MESSAGE_TYPE_RESPONSE,
			Response:    &databaseGrpc.Response{
				ResponseType: databaseGrpc.RESPONSE_TYPE_FIND_RESPONSE,
				FindResponse: &databaseGrpc.FindResponse{
					ResponseStatus: databaseGrpc.RESPONSE_STATUS_SUCCESS,
					Results:        accAny,
					Error:          "",
				},
			},
		}, nil
	}
	return nil, nil
}