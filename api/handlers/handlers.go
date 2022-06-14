package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	omdbChan := make(chan *http.Response)
	vars := mux.Vars(r)
	title := vars["title"]
	go callOmdb(title, omdbChan)
	res := <-omdbChan
	respBody, _ := ioutil.ReadAll(res.Body)

	w.Write(respBody)
}

func callOmdb(title string, c chan *http.Response) {
	omdb_key := os.Getenv("OMDB_API_KEY")
	res, _ := http.Get("https://www.omdbapi.com/?s=" + title + "&page=1&type=movie&apikey=" + omdb_key)
	c <- res
}

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
