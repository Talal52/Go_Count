package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area", "user": username})
}
