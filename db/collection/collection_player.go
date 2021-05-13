package collection

import (
	"context"
	"github.com/TianqiS/database_for_happyball/db"
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
			Keys:    bson.D{{"player_id", 1}},
			Options: (&options.IndexOptions{}).SetUnique(true),
		}
		_, err := indexView.CreateOne(context.TODO(), iModel)
		if err != nil {
			return nil, err
		}
	}
	return playerColl, nil
}

func (playerColl *playerCollection) FindItemsByKey(matchArr []*db.MatchItem) ([]*model.Player, error) {
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

func (playerColl *playerCollection) GetPlayerRankOrderByScore() ([]*model.Player, error) {
	sortState := bson.D{
		{
			"$sort",
			bson.D{
				{
					"highest_score",
					1,
				},
			},
		},
	}
	groupState := bson.D{
		{
			"$group",
			bson.D{
				{
					"_id", nil,
				},
				{
					"tableA",
					bson.D{
						{
							"$push",
							"$$ROOT",
						},
					},
				},
			},
		},
	}
	unwindState := bson.D{
		{
			"$unwind",
			bson.M{
				"path":              "$tableA",
				"includeArrayIndex": "arrIndex",
			},
		},
	}
	projectState := bson.D{
		{
			"$project",
			 bson.D{
				 {
					 "_id",
					 0,
				 }, {
					 "player_id",
					 "$tableA.player_id",
				 }, {
					 "account_id",
					 "$tableA.account_id",
				 }, {
					 "highest_score",
					 "$tableA.highest_score",
				 },
				 {
					 "highest_rank",
					 bson.D{
					 	{
					 		"$add",
					 		bson.A {
					 			"$arrIndex",
					 			int64(1),
							},
						},
					 },
				 },
				 {
					 "update_at",
					 "$tableA.update_at",
				 }, {
					 "create_at",
					 "$tableA.create_at",
				 },
			 },
		},
	}
	cursor, err := playerColl.Aggregate(mongo.Pipeline{sortState, groupState, unwindState, projectState})
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
	return results, nil
}
