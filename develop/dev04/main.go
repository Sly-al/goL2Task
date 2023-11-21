package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	setOfWords := make(map[string]bool)
	lexMap := make(map[string][]string)
	for _, word := range words {
		word = strings.ToLower(word)
		if _, ok := setOfWords[word]; ok {
			continue
		}
		letters := []rune(word)
		sort.Slice(letters, func(i, j int) bool {
			return letters[i] < letters[j]
		})
		lexStr := string(letters)
		lexMap[lexStr] = append(lexMap[lexStr], word)
		setOfWords[word] = true
	}
	ansMap := make(map[string][]string)
	for _, slOfWords := range lexMap {
		if len(slOfWords) > 1 {
			ansMap[slOfWords[0]] = slOfWords[1:]
		}
	}
	return ansMap
}

func main() {
	fmt.Printf("%v", findAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}))
}
