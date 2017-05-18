package acronym

import "strings"

const testVersion = 2

func Abbreviate(phrase string) string {
	var output = []string{}
	words := strings.Split(phrase, " ")

	for _, value := range words {
		output = append(output, string(value[0]))
	}
	return strings.ToUpper(strings.Join(output, ""))
}
