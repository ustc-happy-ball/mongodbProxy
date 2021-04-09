package model

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
