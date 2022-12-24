package host

import (
	"log"
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/handlers"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func StartHosting() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/search/{title}", handlers.HandleSeach)
	r.HandleFunc("/movies", handlers.HandleSave).Methods("POST")
	r.HandleFunc("/movies", handlers.HandleGetAll).Methods("GET")

	http.ListenAndServe(":80", r)
}
