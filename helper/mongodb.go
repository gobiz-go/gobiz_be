package helper

import (
	"context"
	"gocroot/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn model.DBInfo) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		return
	}
	db = client.Database(mconn.DBName)
	return
}

func GetRandomDoc[T any](db *mongo.Database, collection string, size uint) (result []T, err error) {
	filter := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}},
	}
	ctx := context.Background()
	cursor, err := db.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return
	}

	err = cursor.All(ctx, &result)

	return
}
