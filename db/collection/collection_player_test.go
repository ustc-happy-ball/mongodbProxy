package collection

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/driven"
	"github.com/TianqiS/database_for_happyball/model"
	"testing"
	"time"
)

func TestInsertPlayer(t *testing.T) {
	driven.InitClient(configs.MongoURI)

	playerColl, err := GetPlayerCollection()
	if err != nil {
		t.Error("插入时发生了错误", err)
		return
	}
	id, err := playerColl.InsertItem(&model.Player{
		PlayerId:     560,
		AccountId:    "607dabcf63a86c32741a20f4",
		HighestScore: 100,
		HighestRank:  5,
		CreateAt:     time.Now().UnixNano(),
		UpdateAt:     time.Now().UnixNano(),
	})
	if err != nil {
		t.Error("插入时发生了错误", err)
		return
	}
	t.Log("获取的ObjectID为", id)
}

func TestFindPlayerByPlayerId(t *testing.T) {
	driven.InitClient(configs.MongoURI)
	playerColl, _ := GetPlayerCollection()
	player, err := playerColl.FindItemsByKey([]*db.MatchItem{
		{
			Key:      "playerid",
			MatchVal: 500,
		},
	})
	if err != nil {
		t.Error("查找时发生错误", err)
	}
	t.Log("查找到的player为", player[0])
}
