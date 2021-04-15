package db

import "github.com/TianqiS/database_for_happyball/configs"

var CollectionMap = map[int32]Collection{
	configs.ItemPlayer: AccountCollection,
}
