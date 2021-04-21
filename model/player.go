package model

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
)

type Player struct {
	PlayerId int32
	AccountId string
	HighestScore int32
	HighestRank int32
	CreateAt int64
	UpdateAt int64
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
