package auth

import (
	"os"
	"strings"
)

var queries map[string]string

func LoadQueries() error {
	data, err := os.ReadFile("internal/auth/query.sql")
	if err != nil {
		return err
	}

	queries = make(map[string]string)
	lines := strings.Split(string(data), "\n")
	var key string
	for _, line := range lines {
		if strings.HasPrefix(line, "-- name:") {
			key = strings.TrimSpace(strings.TrimPrefix(line, "-- name:"))
			continue
		}
		queries[key] += line + "\n"
	}

	return nil
}
