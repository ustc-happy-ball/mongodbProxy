package driven

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoClient struct {
	database *mongo.Database
	client   *mongo.Client
}

var mc *MongoClient

func InitClient(MongoURI string) {
	log.Println("Initializing MongoDB client...")
	cli := GetMgoCli(MongoURI)
	mc = &MongoClient{
		database: cli.Database(configs.DBName),
		client:   cli,
	}
}

func GetCollection(collectionName string, options *options.CollectionOptions) *mongo.Collection {
	return mc.database.Collection(collectionName, options)
}
