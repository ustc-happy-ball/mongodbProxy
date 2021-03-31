package db

import (
	"context"
	"fmt"
	databaseGrpc "github.com/TianqiS/database_for_happyball/grpc"
	grpc2 "google.golang.org/grpc"
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

func (s *server) SendRequest(ctx context.Context, in *databaseGrpc.Request) (*databaseGrpc.Response, error) {
	InitClient()
	if in.FindItemById != nil {
		acc, err := AccountCollection.FindAccount(in.FindItemById.ItemId)
		if err != nil {
			log.Fatal("查找时发生了错误", err)
		}
		fmt.Println("获取的acc为", acc)
		return &databaseGrpc.Response{ResponseType: databaseGrpc.RESPONSE_TYPE_FIND_RESPONSE}, nil
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
