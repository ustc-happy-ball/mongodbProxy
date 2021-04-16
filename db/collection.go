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
	UpdateItemById(objectId string, operation *request.Operation) error
	UpdateItemByKey(matchArr []*request.MatchItem, operation *request.Operation) error
	DeleteItemById(objectId string) error
	DeleteItemByKey(matchArr []*request.MatchItem) error
	GetModel() interface{}
}
