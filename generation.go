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

// Returns the count of the subdivisions when given the len(trianglesInLines)
func subdivisionsFromLinesCount(lines int) int {
	result := math.Log2(float64(2 * lines / 3))
	return int(result)
}

// Returns slice of pointers to slices of Regions.
// All regions will be with 0 type and 0 elevation.
func generateNewRegions(lines []int) []*[]Region {
	result := make([]*[]Region, len(lines))

	for i := 0; i < len(lines); i++ {
		layer := make([]Region, lines[i])
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

func applyElevationTransformation(planet *[]*[]Region, transformation *[][]int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	y := r.Intn(len(*planet))
	x := r.Intn(len(*(*planet)[y]))

	y2 := float64(y)
	if y > len(*planet)/2 {
		y2 = float64(y) / 2
	}

	// between near pole area and equator
	if float64(len(*planet))/6 <= y2 && y2 < float64(len(*planet))/3 {
		*transformation = rotateMatrix45(*transformation)
	}

	// near the pole
	if 0 <= y2 && y2 < float64(len(*planet))/6 {
		*transformation = rotateMatrix90(*transformation)
	}

	for i, i2 := 0, y; i < len(*transformation); i, i2 = i+1, i2+1 {
		for k, k2 := 0, x; k < len((*transformation)[i]); k, k2 = k+1, k2+1 {
			if i2 >= len(*planet) {
				i2 = 0
			}

			if k2 >= len(*(*planet)[i2]) {
				k2 = 0
			}

			(*(*planet)[i2])[k2].Elevation += int16((*transformation)[i][k])
		}
	}
	println()
}

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
