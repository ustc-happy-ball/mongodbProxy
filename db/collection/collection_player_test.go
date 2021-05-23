package collection

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/driven"
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestInsertPlayer(t *testing.T) {
	skipCI(t)
	driven.InitClient(configs.MongoURIForTest)

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
	skipCI(t)
	driven.InitClient(configs.MongoURIForTest)
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

func TestPlayerCollection_GetPlayerRankOrderByScore(t *testing.T) {
	skipCI(t)
	driven.InitClient(configs.MongoURIForTest)
	playerColl, _ := GetPlayerCollection()
	_, err := playerColl.GetPlayerRankOrderByScore()
	if err != nil {
		logrus.Error("error为", err.Error())
	}
}
