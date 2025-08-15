package handlers

import (
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gin-gonic/gin"
)

func HandleGetMovies(c *gin.Context) {
	filter := c.Query("filter")
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	pageSize := c.Query("pageSize")
	if pageSize == "" {
		pageSize = "25"
	}
	movies := repositories.GetOwnedMovies(filter, page, pageSize)
	c.IndentedJSON(http.StatusOK, movies)
}
