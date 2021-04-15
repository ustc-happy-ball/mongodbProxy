package db

import (
	"context"
	"github.com/TianqiS/database_for_happyball/model"
	"log"
)

type accountCollection struct {
	*BaseCollection
}


var AccountCollection = &accountCollection{
	BaseCollection: NewBaseCollection("Account"),
}

//func (accountCollection *accountCollection) FindOneItemById(objectId string) (*model.Account, error) {
//	result, err := accountCollection.BaseCollection.FindOneItemById(objectId)
//	if err != nil {
//		return nil, err
//	}
//	account := &model.Account{}
//	err = result.Decode(account)
//	if err != nil {
//		return nil, err
//	}
//	return account, nil
//}

func (accountColl *accountCollection) FindItemsByKey(key string, value string) ([]interface{}, error) {
	cursor, err := accountColl.BaseCollection.GetCursorOnKeyValue(key, value)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = cursor.Close(context.TODO())
		if err != nil {
			log.Println("在close cursor时发生了错误")
		}
	}()
	var results []interface{}
	for cursor.Next(context.TODO()) {
		account := &model.Account{}
		err = cursor.Decode(account)
		if err != nil {
			return nil, err
		}
		results = append(results, account)
	}
	return results, nil
}

func (accountColl *accountCollection) GetModel() interface{} {
	return model.Account{}
}

//func (accountCollection *accountCollection) FindItemsByKey(key string, value string) ([]*model.Account, error) {
//	cursor, err := accountCollection.BaseCollection.FindItemsByKey(key, value)
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		err := cursor.Close(context.TODO())
//		if err != nil {
//			log.Println("close cursor的时候发生了错误")
//		}
//	}()
//	var result []*model.Account
//	for cursor.Next(context.TODO()) {
//		account := &model.Account{}
//		err = cursor.Decode(account)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, account)
//	}
//	return result, nil
//}



