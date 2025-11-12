package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	trimmed := strings.TrimSpace(text)
	lowercase := strings.ToLower(trimmed)
	result = strings.Split(lowercase, " ")

	return result
}
