package db

import "go.mongodb.org/mongo-driver/mongo"

type Collection interface {
	GetCollection() *mongo.Collection
	InsertItem(item interface{}) (string, error)
	FindOneItemById(objectId string) (*mongo.SingleResult, error)
	FindOneItemByKey(key string, value interface{}) (*mongo.SingleResult, error)
	UpdateItemById(objectId string, newItem interface{}) error
	UpdateItemByKey(key string, value interface{}, newItem interface{}) error
	DeleteItemById(objectId string) error
	DeleteItemByKey(key string, value interface{}) error
}
