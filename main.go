package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Talal52/Go_Count/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	start := time.Now()
	router := gin.Default()

	// Authentication routes
	router.POST("/login", loginHandler)
	router.GET("/protected", authMiddleware(), protectedHandler)

	router.POST("/readFile", authMiddleware(), func(c *gin.Context) {
		// Hardcoded file name
		fileName := "file.txt"

		// Pass the file path to AnalyzeFileContent
		lines, words, vowels, punctuations, spaces, err := utils.AnalyzeFileContent(fileName)
		if err != nil {
			fmt.Println("Error processing file:", err) // Log the error
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
	})

	// Start the server
	fmt.Println("Starting the server at :4000")
	if err := router.Run(":4000"); err != nil {
		fmt.Println("Could not start server:", err)
	}

	fmt.Printf("Execution time: %s\n", time.Since(start))
}

// === HANDLERS ===

func loginHandler(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := createToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func protectedHandler(c *gin.Context) {
	// Username retrieved from context set by authMiddleware
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area", "user": username})
}

// === JWT FUNCTIONS ===

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(secretKey)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := authHeader[7:]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["username"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("username", claims["username"].(string)) // pass the username along
		c.Next()
	}
}
