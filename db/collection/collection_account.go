package collection

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type accountCollection struct {
	*BaseCollection
}

var accountColl *accountCollection

func GetAccountCollection() (*accountCollection, error) {
	if accountColl == nil {
		go logrus.Debug("Getting account collecting...")
		accountColl = &accountCollection{
			BaseCollection: NewBaseCollection("Account"),
		}
		collection := accountColl.GetCollection()
		indexView := collection.Indexes()
		iModel := mongo.IndexModel{
			Keys:    bson.D{{"phone", 1}},
			Options: (&options.IndexOptions{}).SetUnique(true),
		}
		_, err := indexView.CreateOne(context.TODO(), iModel)
		if err != nil {
			return nil, err
		}
	}
	return accountColl, nil
}

func (accountColl *accountCollection) FindItemsByKey(matchArr []*db.MatchItem) ([]*model.Account, error) {
	go logrus.Debug("Finding item by key...")
	cursor, err := accountColl.BaseCollection.GetCursorOnKeyValue(matchArr)
	if err != nil {
		go logrus.Error(err)
		return nil, err
	}
	var results []*model.Account
	for cursor.Next(context.TODO()) {
		account := &model.Account{}
		err = cursor.Decode(account)
		if err != nil {
			return nil, err
		}
		results = append(results, account)
	}
	err = cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (accountColl *accountCollection) InsertItem(item interface{}) (string, error) {
	go logrus.Debug("Inserting item...")
	account := item.(*model.Account)
	account.ID = primitive.NewObjectID()
	objectId, err := accountColl.BaseCollection.InsertItem(account)
	if err != nil {
		return "", err
	}
	return objectId, nil
}

func (accountColl *accountCollection) GetModel() interface{} {
	return &model.Account{}
}
