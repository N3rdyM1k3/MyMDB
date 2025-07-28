package handlers

import (
    "net/http"

    "github.com/N3rdyM1k3/MyMDB/api/repositories"
    "github.com/gin-gonic/gin"
)

func HandleDeleteMovie(c *gin.Context) {
    id := c.Param("id")
    if err := repositories.DeleteMovie(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "Movie