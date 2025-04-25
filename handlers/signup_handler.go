package handlers

import (
	"net/http"

	"github.com/Talal52/Go_Count/db"
	"github.com/Talal52/Go_Count/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *gin.Context) {
	var u models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Insert the user into the database
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = db.DB.Exec(query, u.Username, string(hashedPassword))
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
