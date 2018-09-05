package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Format: /team/:teamnumber
func GinTeamHandler(c *gin.Context) {
	teamNumber := c.Param("teamnumber")

	c.JSON(http.StatusOK, gin.H{"team": teamNumber})

}