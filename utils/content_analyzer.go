package utils

import (
    "fmt"
    "os"

    "github.com/Talal52/Go_Count/cmd"
    "github.com/Talal52/Go_Count/db"
    "github.com/Talal52/Go_Count/models"
)

func AnalyzeFileContent(filePath string, userID int) (int, int, int, int, int, error) {
    var Lines, Words, Vowels, Punctuations, Spaces int

    // Read the content
    content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return 0, 0, 0, 0, 0, err
    }

    // Process the file content in chunks using goroutines
    routines := 4
    channel := make(chan models.Count)
    chunk := len(content) / routines

    for i := 0; i < routines; i++ {
        start := i * chunk
        end := start + chunk
        if i == routines-1 {
            end = len(content)
        }
        go cmd.Count(string(content[start:end]), channel)
    }

    // Collect results from all goroutines
    for i := 0; i < routines; i++ {
        Counts := <-channel
        Lines += Counts.Lines
        Words += Counts.Words
        Vowels += Counts.Vowels
        Punctuations += Counts.Punctuations
        Spaces += Counts.Spaces
    }

    // Store the results in the database
    err = db.StoreResults(userID, filePath, Lines, Words, Vowels, Punctuations, Spaces)
    if err != nil {
        fmt.Println("Error storing results in database:", err)
        return 0, 0, 0, 0, 0, err
    }

    return Lines, Words, Vowels, Punctuations, Spaces, nil
}