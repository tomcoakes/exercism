package hamming

import (
	"strings"
	"fmt"
)

const testVersion = 6

func Distance(dnaOne, dnaTwo string) (int, error) {
	hammingDistance := 0

	if len(dnaOne) != len(dnaTwo) {
		return 0, fmt.Errorf("DNA strings aren't the same length")
	}

	for index, char := range strings.Split(dnaOne, "") {
		if char != strings.Split(dnaTwo, "")[index] {
			hammingDistance ++
		}
	}
	return hammingDistance, nil
}
