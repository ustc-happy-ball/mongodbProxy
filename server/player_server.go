package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/collection"
	"github.com/TianqiS/database_for_happyball/message"
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
)

type PlayerRpcServer struct {
	databaseGrpc.UnimplementedPlayerServiceServer
}

func GetPlayerServer() *PlayerRpcServer {
	return &PlayerRpcServer{}
}

func (*PlayerRpcServer) PlayerFindByPlayerId(ctx context.Context, req *databaseGrpc.PlayerFindByPlayerIdRequest) (*databaseGrpc.PlayerFindByPlayerIdResponse, error) {
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	players, err := playerColl.FindItemsByKey([]*db.MatchItem{
		{
			Key:      "player_id",
			MatchVal: req.PlayerId,
		},
	})
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	var resMsg *databaseGrpc.PlayerFindByPlayerIdResponse
	if len(players) == 0 {
		resMsg = message.NewPlayerFindByPlayerIdResponse(nil)
	} else {
		resMsg = message.NewPlayerFindByPlayerIdResponse(players[0])
	}
	return resMsg, nil
}
func (*PlayerRpcServer) PlayerAdd(ctx context.Context, req *databaseGrpc.PlayerAddRequest) (*databaseGrpc.PlayerAddResponse, error) {
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	newPlayerPb := req.Player
	newPlayer := &model.Player{}
	_ = newPlayer.FromMessage(newPlayerPb)
	_, err = playerColl.InsertItem(newPlayer)
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	return message.NewPlayerAddResponse(), nil
}

func (*PlayerRpcServer) PlayerUpdateHighestScoreByPlayerId(ctx context.Context, req *databaseGrpc.PlayerUpdateHighestScoreByPlayerIdRequest) (
	*databaseGrpc.PlayerUpdateHighestScoreByPlayerIdResponse, error) {
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	playerId := req.PlayerId
	err = playerColl.UpdateItemByKey([]*db.MatchItem{
		{
			Key: "player_id",
			MatchVal: playerId,
		},
	}, &db.Operation{
		Op: "$set",
		Items: []*db.MatchItem{
			{
				Key:      "highest_score",
				MatchVal:  req.HighestScore,
			},
		},
	})
	if err != nil {
		err = ErrorHandler(err)
		return nil, err
	}
	return message.NewPlayerUpdateHighestScoreByPlayerIdResponse(), nil
}
