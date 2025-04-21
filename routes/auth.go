package routes

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func LoginHandler(c *gin.Context) {
    var u struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if u.Username == "Talal" && u.Password == "123456" {
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

func ProtectedHandler(c *gin.Context) {
    username := c.MustGet("username").(string)
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area", "user": username})
}

func AuthMiddleware() gin.HandlerFunc {
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

        c.Set("username", claims["username"].(string))
        c.Next()
    }
}

func createToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(secretKey)
}