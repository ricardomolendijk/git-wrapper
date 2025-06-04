package main

import "strings"

// removeCommentLines removes lines starting with '#' and trims whitespace. Skips empty lines.
func removeCommentLines(content string) string {
	var result []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		result = append(result, trimmed)
	}

	return strings.Join(result, "\n")
}
