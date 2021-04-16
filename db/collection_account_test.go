package db

import (
	"fmt"
	"github.com/TianqiS/database_for_happyball/internal/event/request"
	"github.com/TianqiS/database_for_happyball/model"
	"log"
	"testing"
	"time"
)

func TestInsertAccount(t *testing.T) {
	InitClient()

	id, err := AccountCollection.InsertItem(&model.Account{
		Name:          "song",
		LoginPassword: "tttqitian",
		AccountAvatar: "www.baidu.com",
		Level:         1,
		Delete:        false,
		Region:        "China",
		Phone:         "17376515082",
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
	InitClient()
	acc, err := AccountCollection.FindOneItemById("605b7267be255a7618e38d6a")
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	fmt.Println("获取的acc为", acc)
}

func TestFindAccountByKey(t *testing.T) {
	InitClient()
	acc, err := AccountCollection.FindItemsByKey([]*request.MatchItem{
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
	InitClient()
	err := AccountCollection.DeleteItemById("605b7267be255a7618e38d6a")
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestDeleteAccountByKey(t *testing.T) {
	InitClient()
	err := AccountCollection.DeleteItemByKey([]*request.MatchItem{})
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestUpdateAccountById(t *testing.T) {
	InitClient()
	err := AccountCollection.UpdateItemById(
		"6078605aca07062a614b7c18",
		&request.Operation{
			Op: "$set",
			Items: []*request.MatchItem{
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
	InitClient()
	err := AccountCollection.UpdateItemByKey(
		[]*request.MatchItem{
			{
				Key: "name",
				MatchVal: "123",
			},
		},
		&request.Operation{
			Op: "$set",
			Items: []*request.MatchItem {
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
