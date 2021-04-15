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
	_, err := AccountCollection.FindItemsByKey([]*request.MatchItem{})
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	//fmt.Println("获取的acc为", acc)
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
	err := AccountCollection.DeleteItemByKey("level", 1)
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}
