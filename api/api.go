package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	omdb_key := os.Getenv("OMDB_API_KEY")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/search/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		res, err := http.Get("https://www.omdbapi.com/?s=" + vars["title"] + "&page=1&type=movie&apikey=" + omdb_key)
		if err != nil {
			fmt.Print("Error", w)
		}
		respBody, _ := ioutil.ReadAll(res.Body)

		w.Write(respBody)
	})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })

	// "https://www.omdbapi.com/?s="+title+"&page=1&type=movie&apikey=1511e143"

	http.ListenAndServe(":80", r)
}
