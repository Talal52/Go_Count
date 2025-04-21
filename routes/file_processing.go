package routes

import (
    "net/http"

    "github.com/Talal52/Go_Count/db"
    "github.com/Talal52/Go_Count/utils"
    "github.com/gin-gonic/gin"
)

func FileProcessingHandler(c *gin.Context) {
    fileName := "file.txt"
    username := c.MustGet("username").(string)

    lines, words, vowels, punctuations, spaces, err := utils.AnalyzeFileContent(fileName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing file", "details": err.Error()})
        return
    }

    db.StoreResults(username, fileName, lines, words, vowels, punctuations, spaces)

    c.JSON(http.StatusOK, gin.H{
        "lines":        lines,
        "words":        words,
        "vowels":       vowels,
        "punctuations": punctuations,
        "spaces":       spaces,
    })
}