package db

import (
	"context"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

func TestClient(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewDbServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.DbMessage{Id: 1,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{
			FindItemById: &databaseGrpc.FindItemById{
				Item:   databaseGrpc.ITEM_PLAYER,
				ItemId: "605b7267be255a7618e38d6a",
			},
		}}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	accountPb := &databaseGrpc.Account{}
	anypb.UnmarshalTo(r.GetResponse().FindResponse.Results, accountPb, proto.UnmarshalOptions{})
	log.Printf("Greeting: %s", r.GetResponse())
}
