package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gorilla/mux"
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

func HandleSeach_Two(w http.ResponseWriter, r *http.Request) {
	omdbChan := make(chan repositories.MovieCollection)
	vars := mux.Vars(r)
	title := vars["title"]
	go repositories.GetOmdbMovies(title, omdbChan)
	m := <-omdbChan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Movies)

}
