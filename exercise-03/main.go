package main

import (
	"fmt"

	myUtils "github.com/nhtera/exercise-03/utils"
)

func main() {
	// Assignment 01
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := myUtils.CountRectangles(arr)
	fmt.Printf("Number of Rectangles: %v\n", count)

	// Assignment 02
	word := "a123bc34d8ef34"
	count = myUtils.NumDifferentIntegers(word)
	fmt.Printf("Number of different Integers: %v\n", count)
}
