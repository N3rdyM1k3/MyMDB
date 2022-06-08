package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/search/{title}", func(w http.ResponseWriter, r *http.Request) {
		omdbChan := make(chan *http.Response)
		vars := mux.Vars(r)
		title := vars["title"]
		go callOmdb(title, omdbChan)
		res := <-omdbChan
		respBody, _ := ioutil.ReadAll(res.Body)

		w.Write(respBody)
	})
	http.ListenAndServe(":80", r)
}

func callOmdb(title string, c chan *http.Response) {
	omdb_key := os.Getenv("OMDB_API_KEY")
	res, _ := http.Get("https://www.omdbapi.com/?s=" + title + "&page=1&type=movie&apikey=" + omdb_key)
	c <- res
}
