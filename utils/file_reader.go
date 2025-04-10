package utils

import (
	"github.com/Talal52/Go_Count/models"
	"fmt"
	"os"
    "github.com/Talal52/Go_Count/cmd"
)

func ReadFile(filePath string) (int, int, int, int, int, error) {
    var Lines, Words, Vowels, Punctuations, Spaces int

    content, err := os.ReadFile(filePath)
    if err!=nil{
        fmt.Println("Error reading file:", err)
        return 0, 0, 0, 0, 0, err
    }
    routines:= 4
	channel := make(chan models.Count)
	chunk := len(content) / routines
	for i := 0; i < routines; i++ {
		start := i * chunk
		end := start + chunk
		go cmd.Count(string(content[start:end]), channel)
    }
    for i := 0; i < routines; i++ {
		Counts := <-channel
		Lines = Lines + Counts.Lines
		Words = Words + Counts.Words
		Vowels = Vowels + Counts.Vowels
		Punctuations = Punctuations + Counts.Punctuations
        Spaces= Spaces+Counts.Spaces
	}
		
		return Lines, Words, Vowels, Punctuations, Spaces, nil

}