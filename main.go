package main

import (
    "fmt"
    "time"

    "github.com/Talal52/Go_Count/db"
    "github.com/Talal52/Go_Count/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    db.InitDB()

    router := gin.Default()

    // Authentication routes
    router.POST("/login", routes.LoginHandler)
    router.GET("/protected", routes.AuthMiddleware(), routes.ProtectedHandler)

    // File processing route
    router.POST("/readFile", routes.AuthMiddleware(), routes.FileProcessingHandler)

    // History route
    router.GET("/history", routes.AuthMiddleware(), routes.HistoryHandler)

    // Start the server
    fmt.Println("Starting the server at :4000")
    if err := router.Run(":4000"); err != nil {
        fmt.Println("Could not start server:", err)
    }

    fmt.Printf("Execution time: %s\n", time.Since(time.Now()))
}