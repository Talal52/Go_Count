package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Talal52/Go_Count/utils"
	"github.com/gin-gonic/gin"
)

func FileProcessingHandler(c *gin.Context) {
	// Get the username from the JWT token
	username := c.MustGet("username").(string)

	// Parse the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Save the uploaded file to a temporary location
	tempDir := "./uploads"
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create upload directory"})
		return
	}
	filePath := filepath.Join(tempDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save uploaded file"})
		return
	}

	// Process the file content
	lines, words, vowels, punctuations, spaces, err := utils.AnalyzeFileContent(filePath, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing file", "details": err.Error()})
		return
	}

	// Return the results
	c.JSON(http.StatusOK, gin.H{
		"lines":        lines,
		"words":        words,
		"vowels":       vowels,
		"punctuations": punctuations,
		"spaces":       spaces,
	})
}
