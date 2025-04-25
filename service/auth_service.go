package service

import (
	"errors"

	"github.com/Talal52/Go_Count/db"
	"github.com/Talal52/Go_Count/models"
	"golang.org/x/crypto/bcrypt"
)

func SignupService(user models.User) error {
	// Validate input
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("could not hash password")
	}

	// Store the user in the database
	return db.CreateUser(user.Username, string(hashedPassword))
}
