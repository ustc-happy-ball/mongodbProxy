package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/message"
	"github.com/TianqiS/database_for_happyball/model"
)

type RpcServer struct {
	databaseGrpc.UnimplementedAccountServiceServer
}

func GetServer() *RpcServer {
	return &RpcServer{}
}

func (*RpcServer) AccountFindByPhone(ctx context.Context, req *databaseGrpc.AccountFindByPhoneRequest) (*databaseGrpc.AccountFindByPhoneResponse, error) {
	accounts, err := db.GetAccountCollection().FindItemsByKey([]*db.MatchItem{
		{
			Key: "phone",
			MatchVal: req.Phone,
		},
	})
	var resMsg *databaseGrpc.AccountFindByPhoneResponse
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		resMsg = message.NewAccountFindByPhoneResponse(nil)
	} else {
		resMsg = message.NewAccountFindByPhoneResponse(accounts[0])
	}
	return resMsg, nil
}
func (*RpcServer) AccountAdd(ctx context.Context, req *databaseGrpc.AccountAddRequest) (*databaseGrpc.AccountAddResponse, error) {
	newAccountPb := req.Account
	newAccount := &model.Account{}
	err := newAccount.FromMessage(newAccountPb)
	if err != nil {
		return nil, err
	}
	_, err = db.GetAccountCollection().InsertItem(newAccount)
	if err != nil {
		return nil, err
	}
	return message.NewAccountAddResponse(), nil
}
