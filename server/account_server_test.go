package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:8890"
)

func TestClientAccountFindByPhone(t *testing.T) {
	//// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewAccountServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.AccountFindByPhoneRequest{Phone: "17376515082"}
	r, err := c.AccountFindByPhone(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	accountPb := r.Account
	log.Printf("response: %s", accountPb.String())
}

func TestClientAccountAdd(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewAccountServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.AccountAddRequest{Account: &databaseGrpc.Account{
		PlayerId:      0,
		LoginPassword: "123124",
		Delete:        false,
		Phone:         "17376515081",
		RecentLogin:   0,
		CreateAt:      0,
		UpdateAt:      0,
	}}
	r, err := c.AccountAdd(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("response: %s", r.String())
}

func TestAccountRpcServer_AccountFindPlayerByAccountId(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewAccountServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.AccountFindPlayerByAccountIdRequest{
		AccountId: "60812ca07adbfd45fd7ae3e9",
	}
	r, err := c.AccountFindPlayerByAccountId(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("response: %s", r.String())
}
