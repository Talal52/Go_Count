package routes

import (
    "github.com/Talal52/Go_Count/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    // Authentication routes
    router.POST("/signup", handlers.SignupHandler)
    router.POST("/login", handlers.LoginHandler)

    // File processing route
    router.POST("/readFile", handlers.FileProcessingHandler)

    // History route
    router.GET("/history", handlers.HistoryHandler)

	// Protected route
	router.GET("/protected", handlers.ProtectedHandler)
}