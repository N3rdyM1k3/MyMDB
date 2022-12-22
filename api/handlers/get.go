package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
)

// https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go
type Payload struct {
	data []interface{}
}

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	dataChan := make(chan Payload)
	go getAllMovies(dataChan)
	movies := <-dataChan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies.data)
}

func getAllMovies(c chan Payload) {
	movies := repositories.GetOwnedMovies()
	p := Payload{movies}
	c <- p
}
