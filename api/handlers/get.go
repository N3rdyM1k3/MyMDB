package handlers

import (
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gin-gonic/gin"
)

func HandleGetMovies(c *gin.Context) {
	movies := repositories.GetOwnedMovies()
	c.IndentedJSON(http.StatusOK, movies)
}
