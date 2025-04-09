package main

import (
	"fmt"
	"os"
	"time"
)

type Count struct {
	Lines        int
	Words        int
	Vowels       int
	Punctuations int
	Spaces       int
}

func main() {
	start := time.Now()

	content, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)
	var count Count

	for _, char := range text {
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
	count.Lines++                            // Increment line count for the last line
	// Print results
	fmt.Println("Number of lines:", count.Lines+1) // +1 to count the last line
	fmt.Println("Number of words:", count.Words)
	fmt.Println("Number of vowels:", count.Vowels)
	fmt.Println("Number of punctuations:", count.Punctuations)
	fmt.Println("Number of spaces:", count.Spaces)

	// Execution time
	fmt.Println("Execution time:", time.Since(start))
}
