package main

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/framework/event"
	"github.com/TianqiS/database_for_happyball/internal/handler"
	"github.com/TianqiS/database_for_happyball/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db.InitClient()
	event.EManager.Register(configs.RequestFindById, handler.FindByIdHandler)
	event.EManager.Register(configs.RequestFindByKey, handler.FindByKeyHandler)
	event.EManager.Register(configs.RequestAddRequest, handler.AddHandler)
	event.EManager.Register(configs.RequestDeleteById, handler.DeleteByIdHandler)
	event.EManager.Register(configs.RequestDeleteByKey, handler.DeleteByKeyHandler)
	event.EManager.Register(configs.RequestUpdateById, handler.UpdateByIdHandler)
	event.EManager.Register(configs.RequestUpdateByKey, handler.UpdateByKeyHandler)

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
