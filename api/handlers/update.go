package handlers

import (
	"net/http"

	"github.com/N3rdyM1k3/MyMDB/api/repositories"
	"github.com/gin-gonic/gin"
)

func HandleUpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie interface{}
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	repositories.UpdateMovie(id, movie)
	c.JSON(http.StatusOK, gin.H{"status": "Movie updated"})
}
