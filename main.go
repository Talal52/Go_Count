package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Talal52/Go_Count/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()
	r := gin.Default()

	// POST route
	r.POST("/readFile", func(c *gin.Context) {
		lines, words, vowels, punctuations, spaces, err := utils.ReadFile("file.txt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
			return
		}

		//results as a JSON response
		c.JSON(http.StatusOK, gin.H{
			"lines":        lines,
			"words":        words,
			"vowels":       vowels,
			"punctuations": punctuations,
			"spaces":       spaces,
		})
	})

	// Start server
	r.Run()

	fmt.Printf("Execution time: %s\n", time.Since(start))
}
