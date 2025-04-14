package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
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

	// File processing route
	router.POST("/readFile", authMiddleware(), func(c *gin.Context) {
		// Get the file name from the request body
		var requestBody struct {
			FileName string `json:"file_name"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Open the specified file
		file, err := os.Open(requestBody.FileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		chunkResults := []gin.H{}
		for scanner.Scan() {
			line := scanner.Text()
			lines, words, vowels, punctuations, spaces, err := utils.AnalyzeFileContent(line)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing chunk"})
				return
			}

			// Append the result of the current chunk
			chunkResults = append(chunkResults, gin.H{
				"lines":        lines,
				"words":        words,
				"vowels":       vowels,
				"punctuations": punctuations,
				"spaces":       spaces,
			})
		}

		if err := scanner.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
			return
		}

		// Return results for all chunks
		c.JSON(http.StatusOK, gin.H{"chunk_results": chunkResults})
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
