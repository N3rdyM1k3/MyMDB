package host

import (
	"log"

	"github.com/N3rdyM1k3/MyMDB/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartHosting() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	router := gin.Default()

	router.GET("/movies", handlers.HandleGetMovies)
	router.POST("/movies", handlers.HandleSaveMovies)
	router.PATCH("/movies/:id", handlers.HandleUpdateMovie)
	router.DELETE("/movies/:id", handlers.HandleDeleteMovie)
	router.GET("/search/:title", handlers.HandleSearch)

	router.Run("localhost:8080")
}
