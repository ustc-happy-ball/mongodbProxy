package model

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/proto/databaseGrpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct { // 里面的字段名一定要大写开头
	ID            primitive.ObjectID `bson:"_id"` // 后面的标签要加上
	PlayerId      int32              `bson:"player_id"`
	LoginPassword string             `bson:"login_password"` // 登录密码
	Delete        bool               `bson:"delete"`         // 当前账号是否注销
	Phone         string             `bson:"phone"`          // 电话
	RecentLogin   int64              `bson:"recent_login"`
	CreateAt      int64              `bson:"create_at"`
	UpdateAt      int64              `bson:"update_at"`
}

func (acc *Account) FromMessage(accountPb *databaseGrpc.Account) error {
	if primitive.IsValidObjectID(accountPb.ObjectId) {
		id, err := primitive.ObjectIDFromHex(accountPb.ObjectId)
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
