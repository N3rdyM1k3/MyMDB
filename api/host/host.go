package main

import (
	//"io/ioutil"
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/search/{title}", handlers.HandleSearch)
	r.HandleFunc("/movies", handlers.HandleSave).Methods("POST")
	// r.HandleFunc("/search/{title}", func(w http.ResponseWriter, r *http.Request) {

	// })
	http.ListenAndServe(":80", r)
}
