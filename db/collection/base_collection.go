package collection

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/TianqiS/database_for_happyball/db/driven"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BaseCollection struct {
	collectionName string
}

func NewBaseCollection(collectionName string) *BaseCollection {
	return &BaseCollection{
		collectionName: collectionName,
	}
}

func (baseColl *BaseCollection) GetCollection() *mongo.Collection {
	return driven.Mc.GetCollection(baseColl.collectionName, nil)
}

func (baseColl *BaseCollection) InsertItem(item interface{}) (string, error) {
	log.Println("Inserting item in base collection...")
	collection := baseColl.GetCollection()
	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (baseColl *BaseCollection) FindOneItemById(objectId string) (*mongo.SingleResult, error) {
	collection := baseColl.GetCollection()
	itemObjectId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(context.TODO(), bson.M{"_id": itemObjectId})
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

func (baseColl *BaseCollection) GetCursorOnKeyValue(matchArr []*db.MatchItem) (*mongo.Cursor, error) {
	collection := baseColl.GetCollection()
	findFilter := make(bson.M)
	for _, mItem := range matchArr {
		findFilter[mItem.Key] = mItem.MatchVal
	}
	cursor, err := collection.Find(context.TODO(), findFilter)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (baseColl *BaseCollection) UpdateItemById(objectId string, operation *db.Operation) error {
	collection := baseColl.GetCollection()
	itemObjectId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return err
	}
	update := make(primitive.M)
	op := make(primitive.M)
	for _, opItem := range operation.Items {
		op[opItem.Key] = opItem.MatchVal
	}
	update[operation.Op] = op
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": itemObjectId}, update)
	if err != nil {
		return err
	}
	return nil
}

func (baseColl *BaseCollection) UpdateItemByKey(matchArr []*db.MatchItem, operation *db.Operation) error {
	collection := baseColl.GetCollection()
	updateFilter := make(bson.M)
	for _, mItem := range matchArr {
		updateFilter[mItem.Key] = mItem.MatchVal
	}
	update := make(primitive.M)
	op := make(primitive.M)
	for _, opItem := range operation.Items {
		op[opItem.Key] = opItem.MatchVal
	}
	update[operation.Op] = op
	_, err := collection.UpdateMany(context.TODO(), updateFilter, update)
	if err != nil {
		return err
	}
	return nil
}

func (baseColl *BaseCollection) DeleteItemById(objectId string) error {
	collection := baseColl.GetCollection()
	itemObjectId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": itemObjectId})
	if err != nil {
		return err
	}
	return nil
}

func (baseColl *BaseCollection) DeleteItemByKey(matchArr []*db.MatchItem) error {
	collection := baseColl.GetCollection()
	deleteFilter := make(bson.M)
	for _, mItem := range matchArr {
		deleteFilter[mItem.Key] = mItem.MatchVal
	}
	_, err := collection.DeleteMany(context.TODO(), deleteFilter)
	if err != nil {
		return err
	}
	return nil
}
