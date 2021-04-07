package db

import "go.mongodb.org/mongo-driver/mongo"

type accountCollection struct {
	*BaseCollection
}


var AccountCollection = &accountCollection{
	BaseCollection: NewBaseCollection("Account"),
}

func (accountCollection *accountCollection) InsertItem(item interface{}) (string, error) {
	result, err := accountCollection.BaseCollection.InsertItem(item)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (accountCollection *accountCollection) GetCollection() *mongo.Collection {
	panic("implement me")
}

func (accountCollection *accountCollection) FindOneItemById(objectId string) (*mongo.SingleResult, error) {
	result, err := accountCollection.BaseCollection.FindOneItemById(objectId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (accountCollection *accountCollection) FindOneItemByKey(key string, value interface{}) (*mongo.SingleResult, error) {
	panic("implement me")
}

func (accountCollection *accountCollection) UpdateItemById(objectId string, newItem interface{}) error {
	panic("implement me")
}

func (accountCollection *accountCollection) UpdateItemByKey(key string, value interface{}, newItem interface{}) error {
	panic("implement me")
}

func (accountCollection *accountCollection) DeleteItemById(objectId string) error {
	panic("implement me")
}

func (accountCollection *accountCollection) DeleteItemByKey(key string, value interface{}) error {
	panic("implement me")
}



