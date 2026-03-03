package handlers

import (
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gin-gonic/gin"
)

func HandleSaveMovies(c *gin.Context) {
	var movies []interface{}
	if err := c.BindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	repositories.SaveMovies(movies)
	c.JSON(http.StatusCreated, gin.H{"status": "Movies saved successfully"})
}
