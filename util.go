package main

import (
	"math"
)

func rotateMatrix45(matrix [][]int) [][]int {
	height := len(matrix)
	width := len(matrix[0])

	size := height + width - 1

	result := make([][]int, size)
	for i := 0; i < size; i++ {
		row := make([]int, size)
		result[i] = row
	}

	// the actual rotation of the matrix
	for i := 0; i < height; i++ {
		for k := 0; k < width; k++ {
			result[k+i][k+height-1-i] = matrix[i][k]
		}
	}

	// after rotating, between the diagonal lines of numbers remains 0s,
	// so we assign to them the value of a neighbour
	// (otherwise 0 can be found next to a big number and the terrain will be too hackly)
	for i := 0; i < height-1; i++ {
		for k := 0; k < width-1; k++ {
			result[k+i+1][k+height-1-i] = result[k+i+1][k+height-2-i]
		}
	}

	return result
}

func rotateMatrix90(matrix [][]int) [][]int {
	height := len(matrix)
	width := len(matrix[0])

	result := make([][]int, width)
	for i := 0; i < width; i++ {
		column := make([]int, height)
		result[i] = column
	}

	for i := 0; i < height; i++ {
		for k := 0; k < width; k++ {
			result[k][height-1-i] = matrix[i][k]
		}
	}

	return result
}

// Returns the count of the subdivisions when given the len(trianglesInLines)
func subdivisionsFromLinesCount(lines int) int {
	result := math.Log2(float64(2 * lines / 3))
	return int(result)
}

func pow4(n int) int {
	return int(math.Pow(4, float64(n)))
}

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
