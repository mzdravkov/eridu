package main

import (
	"math"
	"math/rand"
	"time"
)

// Returns how much triangles should have each line of a icoshpere
// (sphere made from recursive subdivisioning of the triangles of an icosahedron).
// The result is a slice where the index is the index of the line and the value
// the number of triangles in that line.
func trianglesInLines(s int) []int {
	segmentSize := int(math.Pow(2, float64(s-1)))
	result := make([]int, 3*segmentSize)

	triangles := 5
	for i := 0; i < segmentSize; i++ {
		result[i] = triangles
		triangles += 10
	}

	triangles = 10 * segmentSize
	for i := segmentSize; i < 2*segmentSize; i++ {
		result[i] = triangles
	}

	triangles = 10*segmentSize - 5
	for i := 2 * segmentSize; i < 3*segmentSize; i++ {
		result[i] = triangles
		triangles -= 10
	}

	return result
}

type Region struct {
	Type      uint8
	Elevation int16
}

// Returns the count of the subdivisions when given the trianglesInLines
func subdivisionsFromLines(lines []int) int {
	result := 2 * len(lines) / 3
	return int(math.Log2(float64(result)))
}

// Returns slice of pointers to slices of Regions.
// All regions will be with 0 type and 0 elevation.
func generateNewRegions(lines []int) []*[]Region {
	subdivisions := subdivisionsFromLines(lines)
	result := make([]*[]Region, subdivisions)

	for i := 0; i < subdivisions; i++ {
		layer := make([]Region, int(20*math.Pow(4, float64(i))))
		for k := 0; k < len(layer); k++ {
			layer[k] = Region{0, 0}
		}
		result[i] = &layer
	}

	return result
}

// Returns a transformation matrix that can be added to a random submatrix of the field
// (given that both have the same size) to produce a transformation in the relief.
func randomElevationTransformation(subdivisions int) [][]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	width := r.Intn(int(math.Pow(4, float64(subdivisions)-1))) + 1
	height := r.Intn(int(math.Pow(4, float64(subdivisions)-1))) + 1

	matrix := make([][]int, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		matrix[i] = row
	}

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			maxByPosition := int(math.Min(float64(row), float64(column))) + 1

			maxByNeighbours := 0
			if column != 0 && row != 0 {
				maxByNeighbours = int(math.Min(float64(matrix[row-1][column]), float64(matrix[row][column-1]))) + 1
			} else if column != 0 {
				maxByNeighbours = matrix[row][column-1] + 1
			} else if row != 0 {
				maxByNeighbours = matrix[row-1][column] + 1
			}

			max := 0
			if maxByNeighbours != 0 {
				max = int(math.Min(float64(maxByNeighbours), float64(maxByPosition)))
			} else {
				max = maxByPosition
			}

			matrix[row][column] = r.Intn(max + 1)
		}
	}

	return matrix
}

// func applyElevationTransformation(map *[]*[]Region, transformation *[][]int) {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	x := r.Intn(int(20*math.Pow(4, float64(subdivisions)-1)))
// 	height := r.Intn(int(math.Pow(4, float64(subdivisions)-1)))
// }

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
