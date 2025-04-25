package routes

import (
    "github.com/Talal52/Go_Count/handlers"
    "github.com/gin-gonic/gin"
)


func SetupRoutes(router *gin.Engine) {
    // Authentication routes
    router.POST("/signup", handlers.SignupHandler)
    router.POST("/login", handlers.LoginHandler)

    // Protected routes
    protected := router.Group("/")
    protected.Use(AuthMiddleware())
    {
        protected.GET("/protected", handlers.ProtectedHandler)
        protected.POST("/readFile", handlers.FileProcessingHandler)
        protected.GET("/history", handlers.HistoryHandler)
    }
}