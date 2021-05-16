package main

import (
	"flag"
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db/driven"
	log2 "github.com/TianqiS/database_for_happyball/log"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
	"github.com/TianqiS/database_for_happyball/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	//DBName string
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	LogLevel string
	LogToFile bool
)

func initParameter() {
	flag.StringVar(&DBUser, "DBUser","","User name of database" )
	flag.StringVar(&DBPassword, "DBPassword", "","Password of database user")
	flag.StringVar(&DBHost,"Host","","IP address of database")
	flag.StringVar(&DBPort, "Port","","Port to connect to database")
	flag.StringVar(&LogLevel, "LogLevel", "debug", "log level")
	flag.BoolVar(&LogToFile, "LogToFile", false, "move log to file flag")

	flag.Parse()
	if DBUser != "" {
		configs.MongoURI = "mongodb://"+ DBUser + ":"+DBPassword + "@"+ DBHost + ":" + DBPort + "/" + configs.DBName
	} else {
		configs.MongoURI = "mongodb://" + DBUser + ":" + DBPort + "/" + configs.DBName
	}
	if LogLevel == "production" {
		configs.LogLevel = "info"
	} else {
		configs.LogLevel = "debug"
	}
	configs.LogToFile = LogToFile
}

func main() {
	initParameter()
	log2.InitLogSystem()
	if configs.MongoURI == "" {
		go logrus.Error("MongoURI is nil")
	}

	driven.InitClient(configs.MongoURI)

	lis, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		go logrus.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerService(s)

	go logrus.Infof("Serving on %v\n",configs.TcpPort)
	go gracefulQuit(s)
	if err := s.Serve(lis); err != nil {
		go logrus.Errorf("failed to serve: %v", err)
	}
}

func registerService(s *grpc.Server) {
	databaseGrpc.RegisterAccountServiceServer(s, server.GetAccountServer())
	databaseGrpc.RegisterPlayerServiceServer(s, server.GetPlayerServer())
}

// quit gRPC call gracefully
func gracefulQuit(s *grpc.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	s.GracefulStop()
	os.Exit(0)
}
