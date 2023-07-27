package utils

func CountRectangles(rectangles [][]int) int {
	count := 0
	for i := 0; i < len(rectangles); i++ {
		for j := 0; j < len(rectangles[0]); j++ {
			if rectangles[i][j] == 1 {
				count++
				// mark all cells in the current rectangle as visited
				markRectangle(rectangles, i, j)
			}
		}
	}
	return count
}

func markRectangle(rectangles [][]int, i, j int) {
	// mark the current cell as visited
	rectangles[i][j] = -1
	// recursively mark all cells in the current row as visited
	for k := j + 1; k < len(rectangles[0]) && rectangles[i][k] == 1; k++ {
		rectangles[i][k] = -1
	}
	// recursively mark all cells in the current column as visited
	for k := i + 1; k < len(rectangles) && rectangles[k][j] == 1; k++ {
		rectangles[k][j] = -1
	}
	// find the bottom-right corner of the current rectangle
	for m := i + 1; m < len(rectangles); m++ {
		for n := j + 1; n < len(rectangles[0]); n++ {
			if rectangles[m][n] != 1 {
				return
			}
			rectangles[m][n] = -1
		}
	}
}
