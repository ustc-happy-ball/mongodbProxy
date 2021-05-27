package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/collection"
	"github.com/TianqiS/database_for_happyball/message"
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
	"github.com/sirupsen/logrus"
)

type AccountRpcServer struct {
	databaseGrpc.UnimplementedAccountServiceServer
}

func GetAccountServer() *AccountRpcServer {
	return &AccountRpcServer{}
}

func (*AccountRpcServer) AccountFindByPhone(ctx context.Context, req *databaseGrpc.AccountFindByPhoneRequest) (*databaseGrpc.AccountFindByPhoneResponse, error) {
	go logrus.Debug("Receiving AccountFindByPhone request, ",req.Phone)
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	go logrus.Debug("Finding account by phone...\n")
	accounts, err := accountColl.FindItemsByKey([]*db.MatchItem{
		{
			Key:      "phone",
			MatchVal: req.Phone,
		},
	})
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	var resMsg *databaseGrpc.AccountFindByPhoneResponse
	if len(accounts) == 0 {
		go logrus.Debug("Fail to find account by phone number\n")
		resMsg = message.NewAccountFindByPhoneResponse(nil)
	} else {
		resMsg = message.NewAccountFindByPhoneResponse(accounts[0])
	}
	return resMsg, nil
}
func (*AccountRpcServer) AccountAdd(ctx context.Context, req *databaseGrpc.AccountAddRequest) (*databaseGrpc.AccountAddResponse, error) {
	go logrus.Debugf("Receiving AccountAdd request, %v\n", req.Account)
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	newAccountPb := req.Account
	newAccount := &model.Account{}
	err = newAccount.FromMessage(newAccountPb)
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}

	go logrus.Debug("Inserting new account...")
	objectId, err := accountColl.InsertItem(newAccount)
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	return message.NewAccountAddResponse(objectId), nil
}

func (*AccountRpcServer) AccountFindPlayerByAccountId(ctx context.Context, req *databaseGrpc.AccountFindPlayerByAccountIdRequest) (*databaseGrpc.AccountFindPlayerByAccountIdResponse, error) {
	go logrus.Debug("正在调用AccountFindPlayerByAccountId接口")
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	mongoObject, err := accountColl.FindOneItemById(req.AccountId)
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	account := &model.Account{}
	err = mongoObject.Decode(account)
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	players, err := playerColl.FindItemsByKey([]*db.MatchItem{
		{
			Key: "player_id",
			MatchVal: account.PlayerId,
		},
	})
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	var resMsg *databaseGrpc.AccountFindPlayerByAccountIdResponse
	if len(players) == 0 {
		resMsg = message.NewAccountFindPlayerByAccountIdResponse(nil)
		go logrus.Debug("查询到的player信息为nil")
	} else {
		resMsg = message.NewAccountFindPlayerByAccountIdResponse(players[0])
		go logrus.Debugf("查询到的player信息为%v", players[0])
	}
	return resMsg, nil
}

func (*AccountRpcServer) AccountGetAccountByPlayerId(ctx context.Context, req *databaseGrpc.AccountGetAccountInfoByPlayerIdRequest) (*databaseGrpc.AccountGetAccountInfoByPlayerIdResponse, error) {
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	go logrus.Debug("Finding account by phone...\n")
	accounts, err := accountColl.FindItemsByKey([]*db.MatchItem{
		{
			Key:      "player_id",
			MatchVal: req.PlayerId,
		},
	})
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	var resMsg *databaseGrpc.AccountGetAccountInfoByPlayerIdResponse
	if len(accounts) == 0 {
		go logrus.Debug("Fail to find account by phone number\n")
		resMsg = message.NewAccountGetAccountByPlayerIdResponse(nil)
	} else {
		resMsg = message.NewAccountGetAccountByPlayerIdResponse(accounts[0])
	}
	return resMsg, nil
}