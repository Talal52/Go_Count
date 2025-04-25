package db

import (
	"errors"
	"fmt"
	"time"
)

func CreateUser(username, password string) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := DB.Exec(query, username, password)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return errors.New("username already exists")
		}
		return errors.New("could not create user")
	}
	return nil
}
func StoreResults(username, fileName string, lines, words, vowels, punctuations, spaces int) {
	query := `
    INSERT INTO results (username, file_name, lines, words, vowels, punctuations, spaces)
    VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := DB.Exec(query, username, fileName, lines, words, vowels, punctuations, spaces)
	if err != nil {
		fmt.Println("Error storing results:", err)
	}
}

func FetchHistory(username string) []map[string]interface{} {
	query := `
    SELECT file_name, lines, words, vowels, punctuations, spaces, created_at
    FROM results
    WHERE username = $1
    ORDER BY created_at DESC`
	rows, err := DB.Query(query, username)
	if err != nil {
		fmt.Println("Error fetching history:", err)
		return nil
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var fileName string
		var lines, words, vowels, punctuations, spaces int
		var createdAt time.Time

		err := rows.Scan(&fileName, &lines, &words, &vowels, &punctuations, &spaces, &createdAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		history = append(history, map[string]interface{}{
			"file_name":    fileName,
			"lines":        lines,
			"words":        words,
			"vowels":       vowels,
			"punctuations": punctuations,
			"spaces":       spaces,
			"created_at":   createdAt,
		})
	}

	return history
}
