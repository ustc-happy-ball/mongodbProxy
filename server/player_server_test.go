package server

import (
	"context"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestClientPlayerFindByPlayerId(t *testing.T) {
	//// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewPlayerServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.PlayerFindByPlayerIdRequest{PlayerId: 550}
	r, err := c.PlayerFindByPlayerId(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	playerPb := r.Player
	log.Printf("response: %s", playerPb.String())
}

func TestClientPlayerAdd(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewPlayerServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.PlayerAddRequest{
		Player: &databaseGrpc.Player{
			PlayerId:     600,
			AccountId:    "607dabcf63a86c32741a20f5",
			HighestScore: 1000,
			HighestRank:  2,
			CreateAt:     time.Now().UnixNano(),
			UpdateAt:     time.Now().UnixNano(),
		},
	}
	_, err = c.PlayerAdd(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("成功了")
}
