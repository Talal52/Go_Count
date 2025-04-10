package main

import (
    "github.com/Talal52/Go_Count/utils"
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    lines, words, vowels, punctuations, spaces, err := utils.ReadFile("file.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Use the returned values
    fmt.Printf("Lines: %d, Words: %d, Vowels: %d, Punctuations: %d, Spaces: %d\n", lines, words, vowels, punctuations, spaces)
    fmt.Printf("Execution time: %s\n", time.Since(start))
}