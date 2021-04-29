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

func NewAccountAddResponse() *databaseGrpc.AccountAddResponse {
	return &databaseGrpc.AccountAddResponse{}
}
