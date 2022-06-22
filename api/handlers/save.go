package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleSave(w http.ResponseWriter, r *http.Request) {
	var movies []interface{}

	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	saveMovies(movies)
}

func saveMovies(movies []interface{}) {
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
	defer client.Disconnect(ctx)
	collection := client.Database("mymdb").Collection("movies")
	res, insertErr := collection.InsertMany(ctx, movies)
	if insertErr != nil {
		log.Print(insertErr)
	}
	log.Print(len(res.InsertedIDs))
}
