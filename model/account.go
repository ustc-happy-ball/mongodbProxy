package model

import databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"

type Account struct { // 里面的字段名一定要大写开头
	Name string
	LoginPassword string // 登录密码
	AccountAvatar string // 头像
	Level int64 // 当前等级
	Delete bool // 当前账号是否注销
	Region string // 用户的地区
	Phone string // 电话
	MaxScore int64 //最大分数
	CreateAt int64
	UpdateAt int64
}

func (account *Account) EncodeToProto(objectId string) *databaseGrpc.Account {
	accountProto := &databaseGrpc.Account{
		ObjectId:      objectId,
		Name:          account.Name,
		LoginPassword: account.LoginPassword,
		Level:         account.Level,
		Delete:        account.Delete,
		Region:        account.Region,
		Phone:         account.Phone,
		CreateAt:      account.CreateAt,
		UpdateAt:      account.UpdateAt,
		AccountAvatar: account.AccountAvatar,
	}
	return accountProto
}

func DecodeFromProto(accountProto *databaseGrpc.Account) *Account {
	account := &Account{
		Name: accountProto.Name,
		LoginPassword: accountProto.LoginPassword,
		AccountAvatar: accountProto.AccountAvatar,
		Level: accountProto.Level,
		Delete: accountProto.Delete,
		Region: accountProto.Region,
		Phone: accountProto.Phone,
		CreateAt: accountProto.CreateAt,
		UpdateAt: accountProto.UpdateAt,
	}
	return account
}
