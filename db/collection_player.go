package db

import (
	"context"
	"github.com/TianqiS/database_for_happyball/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type playerCollection struct {
	*BaseCollection
}

var playerColl *playerCollection

func GetPlayerCollection() (*playerCollection, error) {
	if playerColl == nil {
		playerColl = &playerCollection{
			BaseCollection: NewBaseCollection("Player"),
		}
		collection := playerColl.GetCollection()
		indexView := collection.Indexes()
		iModel := mongo.IndexModel{
			Keys: bson.D{{"player_id",1}},
			Options: (&options.IndexOptions{}).SetUnique(true),
		}
		_, err := indexView.CreateOne(context.TODO(), iModel)
		if err != nil {
			return nil, err
		}
	}
	return playerColl, nil
}

func (playerColl *playerCollection) FindItemsByKey(matchArr []*MatchItem) ([]*model.Player, error) {
	cursor, err := playerColl.BaseCollection.GetCursorOnKeyValue(matchArr)
	if err != nil {
		return nil, err
	}
	var results []*model.Player
	for cursor.Next(context.TODO()) {
		player := &model.Player{}
		err = cursor.Decode(player)
		if err != nil {
			return nil, err
		}
		results = append(results, player)
	}
	err = cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return results, nil
}


