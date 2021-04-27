package main

import (
	"flag"
	"fmt"
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db/driven"
	databaseGrpc "github.com/TianqiS/database_for_happyball/proto"
	"github.com/TianqiS/database_for_happyball/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	//DBName string
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
)

func initDB() {
	//flag.StringVar(&DBName, "DBName", "","Name of database")
	flag.StringVar(&DBUser, "DBUser","","User name of database" )
	flag.StringVar(&DBPassword, "DBPassword", "","Password of database user")
	flag.StringVar(&DBHost,"Host","","IP address of database")
	flag.StringVar(&DBPort, "Port","","Port to connect to database")

	flag.Parse()
	if DBUser != "" {
		configs.MongoURI = "mongodb://"+ DBUser + ":"+DBPassword + "@"+ DBHost + ":" + DBPort + "/" + configs.DBName
	} else {
		configs.MongoURI = "mongodb://" + DBUser + ":" + DBPort + "/" + configs.DBName
	}
}

func main() {
	initDB()
	if configs.MongoURI == "" {
		 log.Fatalf("MongoURI is nil")
	}

	fmt.Println(configs.MongoURI)
	driven.InitClient(configs.MongoURI)

	lis, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerService(s)

	log.Printf("Serving on %v\n",configs.TcpPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerService(s *grpc.Server) {
	databaseGrpc.RegisterAccountServiceServer(s, server.GetAccountServer())
	databaseGrpc.RegisterPlayerServiceServer(s, server.GetPlayerServer())
}
