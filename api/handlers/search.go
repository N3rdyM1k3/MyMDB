package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	jsmerge "github.com/RaveNoX/go-jsonmerge"
	"github.com/gorilla/mux"
)

func HandleSeach(w http.ResponseWriter, r *http.Request) {
	omdbChan := make(chan repositories.MovieCollection)
	vars := mux.Vars(r)
	title := vars["title"]
	go repositories.SearchOwnedMovies(title, omdbChan)
	m := <-omdbChan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Movies)
	jsmerge.Merge(m.Movies, m.Movies)
}
