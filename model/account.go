package model

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct { // 里面的字段名一定要大写开头
	ID primitive.ObjectID `bson:"_id"` // 后面的标签要加上
	PlayerId int32
	LoginPassword string // 登录密码
	Delete bool // 当前账号是否注销
	Phone string // 电话
	RecentLogin int64
	CreateAt int64
	UpdateAt int64
}

func (acc *Account) FromMessage(accountPb *databaseGrpc.Account) error {
	if primitive.IsValidObjectID(accountPb.AccountId) {
		id, err := primitive.ObjectIDFromHex(accountPb.AccountId)
		if err != nil {
			return err
		}
		acc.ID = id
	}
	acc.PlayerId = accountPb.PlayerId
	acc.LoginPassword = accountPb.LoginPassword
	acc.Delete = accountPb.Delete
	acc.Phone = accountPb.Phone
	acc.RecentLogin = accountPb.RecentLogin
	acc.CreateAt = accountPb.CreateAt
	acc.UpdateAt = accountPb.UpdateAt
	return nil
}
