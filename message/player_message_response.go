package message

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/model"
)

func NewPlayerFindByPlayerIdResponse(player *model.Player) *databaseGrpc.PlayerFindByPlayerIdResponse {
	if player != nil {
		return &databaseGrpc.PlayerFindByPlayerIdResponse{
			Player: &databaseGrpc.Player{
				PlayerId:     player.PlayerId,
				AccountId:    player.AccountId,
				HighestScore: player.HighestScore,
				HighestRank:  player.HighestRank,
				CreateAt:     player.CreateAt,
				UpdateAt:     player.UpdateAt,
			},
		}
	}
	return &databaseGrpc.PlayerFindByPlayerIdResponse{}
}

func NewPlayerAddResponse() *databaseGrpc.PlayerAddResponse {
	return &databaseGrpc.PlayerAddResponse{}
}