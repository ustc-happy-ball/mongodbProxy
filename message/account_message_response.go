package message

import (
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
)

func NewAccountFindByPhoneResponse(acc *model.Account) *databaseGrpc.AccountFindByPhoneResponse {
	if acc != nil {
		return &databaseGrpc.AccountFindByPhoneResponse{
			Account: &databaseGrpc.Account{
				ObjectId:      acc.ID.Hex(),
				PlayerId:      acc.PlayerId,
				LoginPassword: acc.LoginPassword,
				Delete:        acc.Delete,
				Phone:         acc.Phone,
				RecentLogin:   acc.RecentLogin,
				CreateAt:      acc.CreateAt,
				UpdateAt:      acc.UpdateAt,
			},
		}
	}
	return &databaseGrpc.AccountFindByPhoneResponse{}
}

func NewAccountAddResponse(objectId string) *databaseGrpc.AccountAddResponse {
	return &databaseGrpc.AccountAddResponse{ObjectId: objectId}
}

func NewAccountFindPlayerByAccountIdResponse(player *model.Player) *databaseGrpc.AccountFindPlayerByAccountIdResponse {
	if player != nil {
		return &databaseGrpc.AccountFindPlayerByAccountIdResponse{
			PlayerInfo: &databaseGrpc.Player{
				PlayerId:     player.PlayerId,
				AccountId:    player.AccountId,
				HighestScore: player.HighestScore,
				HighestRank:  player.HighestRank,
				CreateAt:     player.CreateAt,
				UpdateAt:     player.UpdateAt,
			},
		}
	}
	return &databaseGrpc.AccountFindPlayerByAccountIdResponse{}
}
