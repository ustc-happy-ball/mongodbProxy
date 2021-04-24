package model

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/proto"
)

type Player struct {
	PlayerId     int32  `bson:"player_id"`
	AccountId    string `bson:"account_id"`
	HighestScore int32  `bson:"highest_score"`
	HighestRank  int32  `bson:"highest_rank"`
	CreateAt     int64  `bson:"create_at"`
	UpdateAt     int64  `bson:"update_at"`
}

func (player *Player) FromMessage(playerPb *databaseGrpc.Player) error {
	player.PlayerId = playerPb.PlayerId
	player.AccountId = playerPb.AccountId
	player.HighestRank = playerPb.HighestRank
	player.HighestScore = playerPb.HighestScore
	player.CreateAt = playerPb.CreateAt
	player.UpdateAt = playerPb.UpdateAt
	return nil
}
