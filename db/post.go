package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID         string `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string `bson:"title" json:"title"`
	Content    string `bson:"content" json:"content"`
	Categories string `bson:"categories" json:"categories"`
	Thumb      string `bson:"thumb" json:"thumb"`
	Date       string `bson:"date" json:"date"`
	Videos     string `bson:"videos" json:"videos"`
	Author     string `bson:"author" json:"author"`
}

func CountPosts() (map[string]int64, [][]Post) {
	mongoDB := ConnectDB()
	collection := mongoDB.Database("blog").Collection("Post")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Count().SetHint("_id_")
	counts := make(map[string]int64)
	var results [][]Post
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetLimit(20)
	categories := []string{"news", "movie", "travel", "doc", "web", "sport", "food", "music", "game", "distraction"}
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
			return map[string]int64{}, [][]Post{}
		}
		counts[category] = count
		var pipeline mongo.Pipeline
		if category != "news" {
			pipeline = mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.D{{Key: "categories", Value: bson.D{{Key: "$regex", Value: category}, {Key: "$options", Value: "i"}}}}}},
				bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 20}}}},
			}
		} else {
			pipeline = mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.D{{Key: "categories", Value: bson.D{{Key: "$regex", Value: category}, {Key: "$options", Value: "i"}}}}}},
				bson.D{{Key: "$sort", Value: bson.D{{Key: "date", Value: -1}}}},
				bson.D{{Key: "$limit", Value: 20}},
			}
		}

		cursor, err := collection.Aggregate(ctx, pipeline)
		if err != nil {
			// Handle error
		}
		defer cursor.Close(context.Background())
		var posts []Post
		for cursor.Next(context.Background()) {
			var doc Post
			err := cursor.Decode(&doc)
			if err != nil {
				// Handle error
			}
			posts = append(posts, doc)
		}

		results = append(results, posts)
		if err := cursor.Err(); err != nil {
			// Handle cursor error
		}
	}

	defer mongoDB.Disconnect(ctx)
	return counts, results
}
