package utils

import (
	"testing"
)

func TestNumDifferentIntegers(t *testing.T) {
	testCases := []struct {
		word     string
		expected int
	}{
		{"abc123def456", 2},
		{"12345", 1},
		{"1a2b3c4d5e", 5},
		{"", 0},
	}

	for _, tc := range testCases {
		result := NumDifferentIntegers(tc.word)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d for word '%s'", tc.expected, result, tc.word)
		}
	}
}
