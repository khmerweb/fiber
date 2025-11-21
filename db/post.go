package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountPosts(category string) int64 {
	mongoDB := ConnectDB()
	collection := mongoDB.Database("blog").Collection("Post")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"categories": bson.M{
			"$regex":   category,
			"$options": "i",
		},
	}
	opts := options.Count().SetHint("_id_")

	count, err := collection.CountDocuments(ctx, filter, opts)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer mongoDB.Disconnect(ctx)
	return count
}
