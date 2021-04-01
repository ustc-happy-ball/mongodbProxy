package main

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	databaseGrpc "github.com/TianqiS/database_for_happyball/grpc"
	"github.com/TianqiS/database_for_happyball/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db.InitClient()
	lis, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	databaseGrpc.RegisterDbServiceServer(s, server.GetServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
