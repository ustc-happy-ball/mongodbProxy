package db

import (
	"context"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

func TestClientFindById(t *testing.T) {
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
		MessageCode: databaseGrpc.MESSAGE_CODE_FIND_BY_ID_REQUEST,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{
			FindItemById: &databaseGrpc.FindItemById{
				FindItem:   databaseGrpc.ITEM_PLAYER,
				ItemId: "606fa5d26ce9e63ec70ac6d6",
			},
		}}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//accountPb := &databaseGrpc.Account{}
	//anypb.UnmarshalTo(r.GetResponse().FindResponse.Results, accountPb, proto.UnmarshalOptions{})
	log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientFindByKey(t *testing.T) {
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
		MessageCode: databaseGrpc.MESSAGE_CODE_FIND_BY_KEY_REQUEST,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{
			FindItemByKey: &databaseGrpc.FindItemByKey{
				FindItem:  databaseGrpc.ITEM_PLAYER,
				MatchItem: []*databaseGrpc.MatchItem{
					&databaseGrpc.MatchItem{
						Key:   "name",
						Match: &databaseGrpc.MatchItem_ValString{ValString: "song"},
					},
				},
			},
		}}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientDeleteById(t *testing.T) {
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
		MessageCode: databaseGrpc.MESSAGE_CODE_DELETE_BY_ID_REQUEST,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{
		DeleteItemById: &databaseGrpc.DeleteItemById{
			DeleteItem: databaseGrpc.ITEM_PLAYER,
			ItemId:     "606fa5d26ce9e63ec70ac6d6",
		},
		}}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientDeleteByKey(t *testing.T) {
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
		MessageCode: databaseGrpc.MESSAGE_CODE_DELETE_BY_KEY_REQUEST,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{DeleteItemByKey: &databaseGrpc.DeleteItemByKey{
			DeleteItem: databaseGrpc.ITEM_PLAYER,
			MatchItem:  []*databaseGrpc.MatchItem{
				&databaseGrpc.MatchItem{
					Key:   "name",
					Match: &databaseGrpc.MatchItem_ValString{ValString: "song"},
				},
			},
		}}}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientAddItem(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := databaseGrpc.NewDbServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	item := &databaseGrpc.Account{
		Name:          "tianqi",
		LoginPassword: "hhh",
		Level:         0,
		Delete:        false,
		Region:        "",
		Phone:         "",
		MaxScore:      0,
		CreateAt:      0,
		UpdateAt:      0,
		AccountAvatar: "",
	}
	anyItem, _ := anypb.New(item)
	requestMsg := &databaseGrpc.DbMessage{Id: 1,
		MessageCode: databaseGrpc.MESSAGE_CODE_ADD_REQUEST,
		MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
		Request: &databaseGrpc.Request{
			AddItem: &databaseGrpc.AddItem{
				AddItem: databaseGrpc.ITEM_PLAYER,
				Item: anyItem,
			},
		},
	}
	r, err := c.SendRequest(ctx, requestMsg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
}