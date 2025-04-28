package handlers

import (
    "net/http"
    "strconv"

    "github.com/Talal52/Go_Count/db"
    "github.com/gin-gonic/gin"
)

func HistoryHandler(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Get pagination parameters
    pageParam := c.DefaultQuery("page", "1")
    limitParam := c.DefaultQuery("limit", "10")

    page, err := strconv.Atoi(pageParam)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }

    limit, err := strconv.Atoi(limitParam)
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
        return
    }

    // Calculate offset
    offset := (page - 1) * limit

    // Fetch paginated history
    history, err := db.FetchHistory(userID.(int), limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching history", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "page":    page,
        "limit":   limit,
        "history": history,
    })
}