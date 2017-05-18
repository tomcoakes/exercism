package raindrops

import (
	"strconv"
	"strings"
	"sort"
)

const testVersion = 3
var conversionConfiguration = map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
var factorsToConvert = []int{3, 5, 7}

func Convert(num int) string {
	factorsToConvert := sortFactors()
	factorsOfNum := getFactors(num, factorsToConvert)

	if len(factorsOfNum) == 0 {
		return strconv.Itoa(num)
	}

	var outputSlice = []string{}
	for _, value := range factorsOfNum {
		outputSlice = append(outputSlice, conversionConfiguration[value])
	}
	return strings.Join(outputSlice, "")
}

func getFactors(num int, factors []int) []int {
	var matchedFactors = []int{}
	for key, value := range factors {
		if (num % value) == 0 {
			matchedFactors = append(matchedFactors, factors[key])
		}
	}

	return matchedFactors
}

func sortFactors() []int {
	keys := []int{}
	for k := range conversionConfiguration {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
