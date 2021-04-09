package model

import "github.com/TianqiS/database_for_happyball/internal/event/info"

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

func (account *Account) ToEvent() interface{} {
	return &info.AccountEvent{
		Name:          account.Name,
		LoginPassword: account.LoginPassword,
		AccountAvatar: account.AccountAvatar,
		Level:         account.Level,
		Delete:        account.Delete,
		Region:        account.Region,
		Phone:         account.Phone,
		MaxScore:      account.MaxScore,
		CreateAt:      account.CreateAt,
		UpdateAt:      account.UpdateAt,
	}
}
