package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type error interface {
	Error() string
}

// API Test
func TestAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Test Successful!"})
	}
}
