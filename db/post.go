package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountPosts() []int64 {
	mongoDB := ConnectDB()
	collection := mongoDB.Database("blog").Collection("Post")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Count().SetHint("_id_")
	var counts []int64
	categories := []string{"news", "movie", "travel", "game", "sport", "doc", "food", "music", "distraction"}
	for _, category := range categories {
		filter := bson.M{
			"categories": bson.M{
				"$regex":   category,
				"$options": "i",
			},
		}
		count, err := collection.CountDocuments(ctx, filter, opts)
		if err != nil {
			log.Fatal(err)
			return []int64{}
		}
		counts = append(counts, count)
	}

	defer mongoDB.Disconnect(ctx)
	return counts
}
