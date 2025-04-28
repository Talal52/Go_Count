package db

import (
	"errors"
	"fmt"
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
func StoreResults(userID int, fileName string, lines, words, vowels, punctuations, spaces int) error {
    query := `
    INSERT INTO results (user_id,file_name, lines, words, vowels, punctuations, spaces)
    VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err := DB.Exec(query, userID, fileName, lines, words, vowels, punctuations, spaces)
    if err != nil {
        return fmt.Errorf("could not store results: %w", err)
    }
    return nil
}
func FetchHistory(userID, limit, offset int) ([]map[string]interface{}, error) {
    query := `
    SELECT file_name, lines, words, vowels, punctuations, spaces, created_at
    FROM results
    WHERE user_id = $1
    ORDER BY created_at DESC
    LIMIT $2 OFFSET $3`
    rows, err := DB.Query(query, userID, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("error fetching history: %w", err)
    }
    defer rows.Close()

    var history []map[string]interface{}
    for rows.Next() {
        var fileName string
        var lines, words, vowels, punctuations, spaces int
        var createdAt string

        err := rows.Scan(&fileName, &lines, &words, &vowels, &punctuations, &spaces, &createdAt)
        if err != nil {
            return nil, fmt.Errorf("error scanning row: %w", err)
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

    return history, nil
}

func GetUserIDByUsername(username string) (int, error) {
    var userID int
    query := `SELECT id FROM users WHERE username = $1`
    err := DB.QueryRow(query, username).Scan(&userID)
    if err != nil {
        return 0, fmt.Errorf("could not find user: %w", err)
    }
    return userID, nil
}