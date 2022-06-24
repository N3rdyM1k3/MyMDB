package handlers

import (
	"net/http"
	"os"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go
type Payload struct {
	data []interface{}
}


func HandleGetAll(w http.ResponseWriter, r *http.Request){
	responseChannel := make(chan *http.Response)

}

func getAllMovies(responseChannel chan *http.Response){
	mongo_key := os.Getenv("MONGO_PWD")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mymdb:" + mongo_key + "@mymdb.fscoq.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("mymdb").Collection("movies")
	var movies []interface{}
	err := collection.Find().Decode(&movies)
	if err != nil {
		log.Fatal(err)
	}
	res := http.Response
	res.
	responseChannel <- res
	
}