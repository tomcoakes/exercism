package pangram

import "strings"

const testVersion = 1

func IsPangram(sentence string) bool {
	lettersMap := map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"i": false,
		"j": false,
		"k": false,
		"l": false,
		"m": false,
		"n": false,
		"o": false,
		"p": false,
		"q": false,
		"r": false,
		"s": false,
		"t": false,
		"u": false,
		"v": false,
		"w": false,
		"x": false,
		"y": false,
		"z": false,
	}

	letters := strings.Split(sentence, "")
	// remove duplicates

}
