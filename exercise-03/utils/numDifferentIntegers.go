package utils

import (
	"strings"
)

func NumDifferentIntegers(word string) int {
	numSet := make(map[string]bool)

	num := ""
	for _, c := range word {
		if c >= '0' && c <= '9' {
			num += string(c)
		} else if num != "" {
			numSet[strings.TrimLeft(num, "0")] = true
			num = ""
		}
	}
	if num != "" {
		numSet[strings.TrimLeft(num, "0")] = true
	}

	return len(numSet)
}
