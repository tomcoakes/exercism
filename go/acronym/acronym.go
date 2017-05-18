package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Abbreviate(phrase string) string {
	var firstLetters = []string{}
	for _, word := range splitPhraseAtSpaces(phrase) {
		for _, word := range splitWordsAtHyphen(word) {
			for _, value := range splitWordsAtCapitals(word) {
				firstLetters = append(firstLetters, string(value[0]))
			}
		}
	}
	return strings.ToUpper(strings.Join(firstLetters, ""))
}

func splitPhraseAtSpaces(phrase string) []string {
	return strings.Split(phrase, " ")
}

func splitWordsAtHyphen(word string) []string {
	return strings.Split(word, "-")
}

func splitWordsAtCapitals(word string) []string {
	var tokens = []string{}
	var currentIndex = 1

	if isWordAllUppercase(word) {
		return []string{word}
	}

	for s := word; s != ""; s = s[currentIndex:] {
		currentIndex = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if currentIndex <= 0 {
			currentIndex = len(s)
		}
		tokens = append(tokens, s[:currentIndex])
	}
	return tokens
}

func isWordAllUppercase(word string) bool {
	return strings.IndexFunc(word, unicode.IsLower) == -1
}
