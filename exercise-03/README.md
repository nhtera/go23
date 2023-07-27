## Exercise 03

### Assignment 01 - Find and count number of rectangles in a 2D array

Goal: Find and count number of rectangles in a 2D array.
- Inputs: An array filled with 0s and 1s.
- Outputs: Number of rectangles filled with 1s
- Given that rectangles are separated and do not touch each other but they can touch the boundary of the array. A single element rectangle counts.

Implement countRectangles to return a number of rectangles filled with 1s.
- Each cell is a rectangle '1' or empty '0', return the number of the rectangles on board.
- Each rectangle can be made in the shape of 1*1, 1*n, m*1, or m*n (m rows, n columns)
- There are no adjacent rectangles

### Code example:

```go
package main

import "fmt"

func countRectangles(rectangles [][]int) int {
	// TODO: Implement countRectangle
	return -1
}

func main() {
	arr := [][]int {
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := countRectangles(arr)
	fmt.Printf("%v", count)
}
```

### Assignment 2 - Implement numDifferentIntegers to return a number of different integers in a word

Goal: Count the number of different integers in a String.
- Outputs: Number of different integers
- Given that a string word consists of digits and lowercase English letters, 2 integers are considered different if their decimal representation without any leading zeros are different.

E.g: "a123bc34d8ef34" => 3 (123, 34, 8)
"A1b01c001" => 1 (1)

### Code example:

```go
package main

import "fmt"

// numDifferentIntegers returns a different integers in a word
func numDifferentIntegers(word string) int {
	// TODO: Implement numDifferentIntegers
	return 0
}
func main() {
  word := "a123bc34d8ef34"

  count := numDifferentIntegers(word)
  fmt.Printf("%v", count)
}
```