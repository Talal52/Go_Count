package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Talal52/Go_Count/db"
	"github.com/Talal52/Go_Count/models"
	"github.com/Talal52/Go_Count/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

func LoginHandler(c *gin.Context) {
	var u models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Query the database for the user
	var storedPassword string
	query := `SELECT password FROM users WHERE username = $1`
	err := db.DB.QueryRow(query, u.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(u.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token
	tokenString, err := utils.GenerateToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
