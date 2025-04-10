package cmd

import (
	"github.com/Talal52/Go_Count/models"
	"fmt"
)

func Count(fileContent string, channel chan models.Count) {
	var count models.Count
	for _, char := range fileContent {
		switch {
		case char == ' ':
			count.Spaces++
		case char == '\t' || char == '\r':
			count.Words++
		case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u':
			count.Vowels++
		case char == '.' || char == '!' || char == '?' || char == ',' || char == ':' || char == ';' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
			count.Punctuations++
			if char == ',' {
				count.Words++
			}
			count.Words++
		case char == '\n':
			count.Lines++

		}
	}

	count.Words = count.Lines + count.Spaces
	count.Lines++ // Increment line count for the last line
	// Print results
	// fmt.Println("Number of lines:", count.Lines+1) // +1 to count the last line
	// fmt.Println("Number of words:", count.Words)
	// fmt.Println("Number of vowels:", count.Vowels)
	// fmt.Println("Number of punctuations:", count.Punctuations)
	// fmt.Println("Number of spaces:", count.Spaces)
	fmt.Printf("chunk results=====Lines:%d, Words:%d, Vowels:%d, Punctuations:%d, Spaces:%d\n", count.Lines, count.Words, count.Vowels, count.Punctuations, count.Spaces)
	channel <- count

}
