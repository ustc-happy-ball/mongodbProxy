package db

import (
	"context"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address     = "localhost:50051"
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
	r, err := c.SendRequest(ctx, &databaseGrpc.Request{FindItemById: &databaseGrpc.FindItemById{
		Item:   databaseGrpc.ITEM_PLAYER,
		ItemId: "605b7267be255a7618e38d6a",
	}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponseType())
}

