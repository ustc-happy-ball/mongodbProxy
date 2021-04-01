package db

import (
	"context"
	"github.com/TianqiS/database_for_happyball/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountCollection struct {}


var AccountCollection = &accountCollection{}

func (this *accountCollection) getCollection() *mongo.Collection {
	return Mc.GetCollection("Account", nil)
}

func (this *accountCollection) InsertAccount(account *model.Account) (string, error) {
	collection := this.getCollection()
	insertResult, err := collection.InsertOne(context.TODO(), account)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).String(), nil
}

func (this *accountCollection) FindAccount(accountId string) (*model.Account, error) {
	collection := this.getCollection()
	accIdObject, err := primitive.ObjectIDFromHex(accountId)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(context.TODO(), bson.M{"_id": accIdObject})
	if result.Err() != nil {
		return nil, result.Err()
	}
	findAcc := &model.Account{}
	err = result.Decode(findAcc)
	if err != nil {
		return nil, err
	}
	return findAcc, nil
}

func (this *accountCollection) UpdateAccount(accountId string, account *model.Account) {
	collection := this.getCollection()
	accIdObject, err := primitive.ObjectIDFromHex(accountId)
	update := bson.M{
		"$set": bson.M{
			"phone": account.Phone,
		},
	}

	if err != nil {
		return
	}
	collection.UpdateByID(context.TODO(), accIdObject, update)
}