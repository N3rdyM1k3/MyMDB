package handlers

import (
	"api/repositories"
	"encoding/json"
	"net/http"
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
	movies := repositories.GetMovies()
	p := Payload{movies}
	c <- p
}
