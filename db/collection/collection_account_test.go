package collection

import (
	"fmt"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/driven"
	"github.com/TianqiS/database_for_happyball/model"
	"log"
	"testing"
	"time"
)

func TestInsertAccount(t *testing.T) {
	driven.InitClient()

	accountColl, _ := GetAccountCollection()
	id, err := accountColl.InsertItem(&model.Account{
		LoginPassword: "tttqitian",
		Delete:        false,
		Phone:         "17376515083",
		CreateAt:      time.Now().UnixNano(),
		UpdateAt:      0,
	})
	if err != nil {
		log.Println("插入时发生了错误", err)
		return
	}
	log.Println("获取的ID为", id)
}

func TestFindAccountById(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	acc, err := accountColl.FindOneItemById("605b7267be255a7618e38d6a")
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	fmt.Println("获取的acc为", acc)
}

func TestFindAccountByKey(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	acc, err := accountColl.FindItemsByKey([]*db.MatchItem{
		{
			Key: "name",
			MatchVal: "tianqi",
		},
	})
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	fmt.Println("获取的acc为", acc)
}

func TestDeleteAccountById(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	err := accountColl.DeleteItemById("605b7267be255a7618e38d6a")
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestDeleteAccountByKey(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	err := accountColl.DeleteItemByKey([]*db.MatchItem{})
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestUpdateAccountById(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	err := accountColl.UpdateItemById(
		"6078605aca07062a614b7c18",
		&db.Operation{
			Op: "$set",
			Items: []*db.MatchItem{
				{
					Key: "name",
					MatchVal: "qiqi",
				},
			},
		},
		)
	if err != nil {
		t.Error("更新时发生了错误", err)
	}
}

func TestUpdateAccountByKey(t *testing.T) {
	driven.InitClient()
	accountColl, _ := GetAccountCollection()
	err := accountColl.UpdateItemByKey(
		[]*db.MatchItem{
			{
				Key: "name",
				MatchVal: "123",
			},
		},
		&db.Operation{
			Op: "$set",
			Items: []*db.MatchItem{
				{
					Key: "name",
					MatchVal: "12345",
				},
				{
					Key: "level",
					MatchVal: int64(3),
				},
			},
		},
		)
	if err != nil {
		t.Error("更新时发生了错误", err)
	}
}
