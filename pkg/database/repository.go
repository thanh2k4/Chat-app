package auth

import (
	"os"
	"strings"
)

type Query struct {
	SQL  string
	Type string
}

func LoadQueries(path string) (map[string]Query, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	queries := make(map[string]Query)
	lines := strings.Split(string(data), "\n")
	var key, queryType string
	var builder strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "-- name:") {
			if key != "" {
				queries[key] = Query{SQL: builder.String(), Type: queryType}
			}

			parts := strings.Fields(strings.TrimPrefix(line, "-- name:"))
			key = parts[0]
			queryType = ""
			if len(parts) > 1 {
				queryType = parts[1]
			}

			builder.Reset()
			continue
		}

		if key != "" {
			builder.WriteString(line + "\n")
		}
	}

	if key != "" {
		queries[key] = Query{SQL: builder.String(), Type: queryType}
	}

	return queries, nil
}
