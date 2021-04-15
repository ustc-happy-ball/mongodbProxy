package db

import (
	"github.com/TianqiS/database_for_happyball/internal/event/request"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection interface {
	GetCollection() *mongo.Collection
	InsertItem(item interface{}) (string, error)
	FindOneItemById(objectId string) (*mongo.SingleResult, error)
	FindItemsByKey(matchArr []*request.MatchItem) ([]interface{}, error)
	UpdateItemById(objectId string, newItem interface{}) error
	UpdateItemByKey(key string, value interface{}, newItem interface{}) error
	DeleteItemById(objectId string) error
	DeleteItemByKey(matchArr []*request.MatchItem) error
	GetModel() interface{}
}
