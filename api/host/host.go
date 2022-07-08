package host

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func StartHosting() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/search/{title}", handlers.HandleSearch)
	r.HandleFunc("/movies", handlers.HandleSave).Methods("POST")
	r.HandleFunc("/movies", handlers.HandleGetAll).Methods("GET")

	http.ListenAndServe(":80", r)
}
