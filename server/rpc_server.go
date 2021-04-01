package server

import (
	"context"
	"fmt"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
	"log"
)

type RpcServer struct {
	databaseGrpc.UnimplementedDbServiceServer
}

func GetServer() *RpcServer {
	return &RpcServer{}
}

func (s *RpcServer) SendRequest(ctx context.Context, in *databaseGrpc.Request) (*databaseGrpc.Response, error) {
	if in.FindItemById != nil {
		acc, err := db.AccountCollection.FindAccount(in.FindItemById.ItemId)
		if err != nil {
			log.Fatal("查找时发生了错误", err)
		}
		fmt.Println("获取的acc为", acc)
		return &databaseGrpc.Response{ResponseType: databaseGrpc.RESPONSE_TYPE_FIND_RESPONSE}, nil
	}
	return nil, nil
}