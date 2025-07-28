package handlers

import (
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gin-gonic/gin"
)

func HandleSearch(c *gin.Context) {
	title := c.Param("title")
	omdbChan := make(chan repositories.MovieCollection)
	go repositories.SearchOwnedMovies(title, omdbChan)
	m := <-omdbChan
	c.IndentedJSON(http.StatusOK, m.Movies)
}
