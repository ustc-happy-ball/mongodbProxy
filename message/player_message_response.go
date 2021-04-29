package message

import (
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
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

func NewPlayerUpdateHighestScoreByPlayerIdResponse() *databaseGrpc.PlayerUpdateHighestScoreByPlayerIdResponse {
	return &databaseGrpc.PlayerUpdateHighestScoreByPlayerIdResponse{}
}