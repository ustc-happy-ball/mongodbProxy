package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/collection"
	"github.com/TianqiS/database_for_happyball/message"
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/TianqiS/database_for_happyball/proto"
	"log"
)

type AccountRpcServer struct {
	databaseGrpc.UnimplementedAccountServiceServer
}

func GetAccountServer() *AccountRpcServer {
	return &AccountRpcServer{}
}

func (*AccountRpcServer) AccountFindByPhone(ctx context.Context, req *databaseGrpc.AccountFindByPhoneRequest) (*databaseGrpc.AccountFindByPhoneResponse, error) {
	log.Printf("Receiving AccountFindByPhone request, %v\n",req.Phone)
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		return nil, err
	}
	accounts, err := accountColl.FindItemsByKey([]*db.MatchItem{
		{
			Key:      "phone",
			MatchVal: req.Phone,
		},
	})
	if err != nil {
		return nil, err
	}
	var resMsg *databaseGrpc.AccountFindByPhoneResponse
	if len(accounts) == 0 {
		resMsg = message.NewAccountFindByPhoneResponse(nil)
	} else {
		resMsg = message.NewAccountFindByPhoneResponse(accounts[0])
	}
	return resMsg, nil
}
func (*AccountRpcServer) AccountAdd(ctx context.Context, req *databaseGrpc.AccountAddRequest) (*databaseGrpc.AccountAddResponse, error) {
	log.Printf("Receiving AccountAdd request, %v\n", req.Account)
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		return nil, err
	}
	newAccountPb := req.Account
	newAccount := &model.Account{}
	err = newAccount.FromMessage(newAccountPb)
	if err != nil {
		return nil, err
	}
	_, err = accountColl.InsertItem(newAccount)
	if err != nil {
		return nil, err
	}
	return message.NewAccountAddResponse(), nil
}
