package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

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
