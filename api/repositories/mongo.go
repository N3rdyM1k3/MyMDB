package repositories

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchOwnedMovies(title string, c chan MovieCollection) {
	client, ctx := buildClient()
	defer client.Disconnect(ctx)
	collection := client.Database("mymdb").Collection("movies")
	var movies []interface{}
	cursor, err := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		for cursor.Next(ctx) {
			var m bson.M
			e := cursor.Decode(&m)
			if e != nil {
				log.Fatal(e)
			} else {
				movies = append(movies, m)
			}
		}
	}
	var m MovieCollection
	m.Movies = movies
	c <- m
}

func GetOwnedMovies() []interface{} {
	client, ctx := buildClient()
	defer client.Disconnect(ctx)
	collection := client.Database("mymdb").Collection("movies")
	var movies []interface{}
	cursor, err := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		for cursor.Next(ctx) {
			var m bson.M
			e := cursor.Decode(&m)
			if e != nil {
				log.Fatal(e)
			} else {
				movies = append(movies, m)
			}
		}
	}
	return movies

}

func SaveOwnedMovies(movies []interface{}) {
	client, ctx := buildClient()
	defer client.Disconnect(ctx)
	collection := client.Database("mymdb").Collection("movies")
	res, insertErr := collection.InsertMany(ctx, movies)
	if insertErr != nil {
		log.Print(insertErr)
	}
	log.Print(len(res.InsertedIDs))
}

func buildClient() (*mongo.Client, context.Context) {
	mongo_key := os.Getenv("MONGO_PWD")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mymdb:" + mongo_key + "@mymdb.fscoq.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}
