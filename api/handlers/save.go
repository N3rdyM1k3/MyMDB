package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
)

func HandleSave(w http.ResponseWriter, r *http.Request) {
	var movies []interface{}

	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repositories.SaveOwnedMovies(movies)
}
