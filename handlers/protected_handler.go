package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ProtectedHandler(c *gin.Context) {
    username, exists := c.Get("username")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to the protected area",
        "user":    username,
    })
}