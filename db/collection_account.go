package db

import (
	"context"
	"github.com/TianqiS/database_for_happyball/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type accountCollection struct {
	*BaseCollection
}

func init() {

}

var accCollection *accountCollection

func GetAccountCollection() *accountCollection {
	if accCollection == nil {
		accCollection =  &accountCollection{
			BaseCollection: NewBaseCollection("Account"),
		}
		collection := accCollection.GetCollection()
		indexView := collection.Indexes()
		iModel := mongo.IndexModel{
			Keys: "phone",
			Options: (&options.IndexOptions{}).SetUnique(true),
		}
		indexView.CreateOne(context.TODO(), iModel)
	}
	return accCollection
}

func (accountColl *accountCollection) FindItemsByKey(matchArr []*MatchItem) ([]*model.Account, error) {
	cursor, err := accountColl.BaseCollection.GetCursorOnKeyValue(matchArr)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = cursor.Close(context.TODO())
		if err != nil {
			log.Println("在close cursor时发生了错误")
		}
	}()
	var results []*model.Account
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
	return &model.Account{}
}



