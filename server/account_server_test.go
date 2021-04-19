package server

import (
	"context"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
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

func TestClientFindByKey(t *testing.T) {
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
		Phone:         "17376515082",
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

func TestClientDeleteById(t *testing.T) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := databaseGrpc.NewDbServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//requestMsg := &databaseGrpc.DbMessage{Id: 1,
	//	MessageCode: databaseGrpc.MESSAGE_CODE_DELETE_BY_ID_REQUEST,
	//	MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
	//	Request: &databaseGrpc.Request{
	//	DeleteItemById: &databaseGrpc.DeleteItemById{
	//		DeleteItem: databaseGrpc.ITEM_PLAYER,
	//		ItemId:     "606fa5d26ce9e63ec70ac6d6",
	//	},
	//	}}
	//r, err := c.SendRequest(ctx, requestMsg)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientDeleteByKey(t *testing.T) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := databaseGrpc.NewDbServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//requestMsg := &databaseGrpc.DbMessage{Id: 1,
	//	MessageCode: databaseGrpc.MESSAGE_CODE_DELETE_BY_KEY_REQUEST,
	//	MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
	//	Request: &databaseGrpc.Request{DeleteItemByKey: &databaseGrpc.DeleteItemByKey{
	//		DeleteItem: databaseGrpc.ITEM_PLAYER,
	//		MatchItem:  []*databaseGrpc.MatchItem{
	//			&databaseGrpc.MatchItem{
	//				Key:   "name",
	//				MatchVal: &databaseGrpc.MatchItem_ValString{ValString: "song"},
	//			},
	//		},
	//	}}}
	//r, err := c.SendRequest(ctx, requestMsg)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientAddItem(t *testing.T) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := databaseGrpc.NewDbServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//item := &databaseGrpc.Account{
	//	Name:          "tianqi",
	//	LoginPassword: "hhh",
	//	Level:         5,
	//	Delete:        false,
	//	Region:        "中国",
	//	Phone:         "17376515082",
	//	MaxScore:      5,
	//	CreateAt:      0,
	//	UpdateAt:      0,
	//	AccountAvatar: "",
	//}
	//anyItem, _ := anypb.New(item)
	//requestMsg := &databaseGrpc.DbMessage{Id: 1,
	//	MessageCode: databaseGrpc.MESSAGE_CODE_ADD_REQUEST,
	//	MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
	//	Request: &databaseGrpc.Request{
	//		AddItem: &databaseGrpc.AddItem{
	//			AddItem: databaseGrpc.ITEM_PLAYER,
	//			Item: anyItem,
	//		},
	//	},
	//}
	//r, err := c.SendRequest(ctx, requestMsg)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientUpdateById(t *testing.T) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := databaseGrpc.NewDbServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//requestMsg := &databaseGrpc.DbMessage{Id: 1,
	//	MessageCode: databaseGrpc.MESSAGE_CODE_UPDATE_BY_ID_REQUEST,
	//	MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
	//	Request: &databaseGrpc.Request{
	//		UpdateItemById: &databaseGrpc.UpdateItemById{
	//			ObjectId:   "6078605aca07062a614b7c18",
	//			UpdateItem: databaseGrpc.ITEM_PLAYER,
	//			Operation:  &databaseGrpc.Operation{
	//				Op:    "$set",
	//				Items: []*databaseGrpc.MatchItem {
	//					&databaseGrpc.MatchItem{
	//						Key:      "name",
	//						MatchVal: &databaseGrpc.MatchItem_ValString{ValString: "tianqi"},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//r, err := c.SendRequest(ctx, requestMsg)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetResponse())
}

func TestClientUpdateByKey(t *testing.T) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := databaseGrpc.NewDbServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//requestMsg := &databaseGrpc.DbMessage{Id: 1,
	//	MessageCode: databaseGrpc.MESSAGE_CODE_UPDATE_BY_KEY_REQUEST,
	//	MessageType: databaseGrpc.MESSAGE_TYPE_REQUEST,
	//	Request: &databaseGrpc.Request{
	//		UpdateItemByKey: &databaseGrpc.UpdateItemByKey{
	//			MatchItem: []*databaseGrpc.MatchItem {
	//				{
	//					Key: "name",
	//					MatchVal: &databaseGrpc.MatchItem_ValString{ValString: "12345"},
	//				},
	//			},
	//			UpdateItem: databaseGrpc.ITEM_PLAYER,
	//			Operation:  &databaseGrpc.Operation{
	//				Op:    "$set",
	//				Items: []*databaseGrpc.MatchItem {
	//					&databaseGrpc.MatchItem{
	//						Key:      "name",
	//						MatchVal: &databaseGrpc.MatchItem_ValString{ValString: "tianqi"},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//r, err := c.SendRequest(ctx, requestMsg)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetResponse())
}