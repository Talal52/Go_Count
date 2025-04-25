package handlers

import (
    "net/http"

    "github.com/Talal52/Go_Count/db"
    "github.com/gin-gonic/gin"
)

func HistoryHandler(c *gin.Context) {
    username := c.MustGet("username").(string)
    history := db.FetchHistory(username)
    c.JSON(http.StatusOK, gin.H{"history": history})
}