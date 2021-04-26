package driven

import (
	"context"
	"github.com/TianqiS/database_for_happyball/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mgoCli *mongo.Client

func initEngine(MongoURI string) {
	var err error
	log.Println("Initializing MongoDB engine...")
	clientOptions := options.Client().ApplyURI(MongoURI)
	clientOptions.SetMaxPoolSize(configs.MongoPoolSize) // 连接池配置
	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMgoCli(MongoURI string) *mongo.Client {
	if mgoCli == nil {
		initEngine(MongoURI)
	}
	return mgoCli
}
