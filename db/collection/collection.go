package collection

import (
	"github.com/TianqiS/database_for_happyball/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection interface {
	GetCollection() *mongo.Collection
	InsertItem(item interface{}) (string, error)
	FindOneItemById(objectId string) (*mongo.SingleResult, error)
	UpdateItemById(objectId string, operation *db.Operation) error
	UpdateItemByKey(matchArr []*db.MatchItem, operation *db.Operation) error
	DeleteItemById(objectId string) error
	DeleteItemByKey(matchArr []*db.MatchItem) error
	GetModel() interface{}
}
