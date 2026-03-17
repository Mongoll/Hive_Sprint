package main

import (
	"fmt"
	"strings"
)

func splitIntoWords(s string) []string {
	return strings.Fields(s)
}

func firstUniqueWord(s string) string {
	wordCount := make(map[string]int)
	words := splitIntoWords(s)

	for _, word := range words {
		wordCount[word]++
	}

	for _, word := range words {
		if wordCount[word] == 1 {
			return word
		}
	}

	return "none"
}

func main() {
	text := "hello hello world world"
	result := firstUniqueWord(text)
	fmt.Println(result)
}