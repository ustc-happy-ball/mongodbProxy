package db

import (
	"context"
	"fmt"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/model"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net"
	"testing"
)

type server struct {
	databaseGrpc.UnimplementedDbServiceServer
}

const (
	port = ":50051"
)

func (s *server) SendRequest(ctx context.Context, in *databaseGrpc.DbMessage) (*databaseGrpc.DbMessage, error) {
	if in.GetRequest().GetFindItemById() != nil {
		obj, err := AccountCollection.FindOneItemById(in.GetRequest().GetFindItemById().ItemId)
		acc := model.Account{}
		obj.Decode(acc)
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
				ResponseStatus: databaseGrpc.RESPONSE_STATUS_OK,
				Error: "",
				FindResponse: &databaseGrpc.FindResponse{
					Results:        accAny,
				},
			},
		}, nil
	}
	return nil, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc2.NewServer()
	databaseGrpc.RegisterDbServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
