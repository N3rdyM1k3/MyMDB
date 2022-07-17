package handlers

import (
	"api/repositories"
	"encoding/json"
	"net/http"
)

func HandleSave(w http.ResponseWriter, r *http.Request) {
	var movies []interface{}

	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repositories.SaveMovies(movies)
}
