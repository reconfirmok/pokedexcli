package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	wordList := strings.Fields(lowerText)
	return wordList
}
