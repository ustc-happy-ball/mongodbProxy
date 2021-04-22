package main

import (
	"github.com/TianqiS/database_for_happyball/configs"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db/driven"
	"github.com/TianqiS/database_for_happyball/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	driven.InitClient()

	lis, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerService(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerService(s *grpc.Server) {
	databaseGrpc.RegisterAccountServiceServer(s, server.GetAccountServer())
	databaseGrpc.RegisterPlayerServiceServer(s, server.GetPlayerServer())
}
