package utils

import "testing"

func TestCountRectangles(t *testing.T) {
	testCases := []struct {
		rectangles [][]int
		expected   int
	}{
		{
			rectangles: [][]int{
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: 6,
		},
		{
			rectangles: [][]int{
				{1, 0, 0, 1, 1, 0, 1},
				{0, 0, 0, 1, 1, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{1, 1, 0, 1, 1, 0, 0},
				{1, 0, 0, 1, 1, 0, 0},
				{1, 0, 0, 0, 0, 0, 1},
			},
			expected: 9,
		},
		{
			rectangles: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			expected: 1,
		},
		{
			rectangles: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		result := CountRectangles(tc.rectangles)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, result)
		}
	}
}
