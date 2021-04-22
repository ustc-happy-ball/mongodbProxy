package main

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
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
	databaseGrpc.RegisterAccountServiceServer(s, server.GetAccountServer())
	databaseGrpc.RegisterPlayerServiceServer(s, server.GetPlayerServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
