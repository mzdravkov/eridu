package main

import (
	"math/rand"
	"time"
)

type Region struct {
	Type      uint8
	Elevation int16
}

// Returns how much triangles should have each line of a icoshpere
// (sphere made from recursive subdivisioning of the triangles of an icosahedron).
// The result is a slice where the index is the index of the line and the value
// the number of triangles in that line.
func trianglesInLines(s int) []int {
	segmentSize := pow2(s - 1)
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
	width := r.Intn(pow4(subdivisions-1)/2) + 1
	height := r.Intn(pow4(subdivisions-1)/2) + 1

	matrix := make([][]int, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		matrix[i] = row
	}

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			maxByPosition := min(row, column) + 1

			maxByNeighbours := 0
			if column != 0 && row != 0 {
				maxByNeighbours = min(matrix[row-1][column], matrix[row][column-1]) + 1
			} else if column != 0 {
				maxByNeighbours = matrix[row][column-1] + 1
			} else if row != 0 {
				maxByNeighbours = matrix[row-1][column] + 1
			}

			max := 0
			if maxByNeighbours != 0 {
				max = min(maxByNeighbours, maxByPosition)
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

	rows := float64(len(*planet))

	y := r.Intn(len(*planet))
	x := r.Intn(len(*(*planet)[y]))

	y2 := float64(y)
	if y > len(*planet)/2 {
		y2 /= 2
	}

	// between near pole area and equator
	if rows/6 <= y2 && y2 < rows/3 {
		*transformation = rotateMatrix45(*transformation)
	}

	// near the pole
	if 0 <= y2 && y2 < rows/6 {
		*transformation = rotateMatrix90(*transformation)
	}

	sign := 1
	if r.Intn(100)%2 == 0 {
		sign = -1
	}

	for i, i2 := 0, y; i < len(*transformation); i, i2 = i+1, i2+1 {
		for k, k2 := 0, x; k < len((*transformation)[i]); k, k2 = k+1, k2+1 {
			if i2 >= len(*planet) {
				i2 = 0
			}

			if k2 >= len(*(*planet)[i2]) {
				k2 = 0
			}

			(*(*planet)[i2])[k2].Elevation += int16((*transformation)[i][k] * sign)
		}
	}
}

func transformPlanetRelief(planet *[]*[]Region) {
	// the maximum size of the transformation matrix is 20 times less than the planet size
	// so the average transformation matrix size is 40 times less than the planet size.
	// so i < 40 will have most of the areas transformed a bit, but i < 80 will increase the change of
	// not having boring non-transformed areas
	for i := 0; i < 40; i++ {
		subdivisions := subdivisionsFromLinesCount(len(*planet))
		transformation := randomElevationTransformation(subdivisions)

		applyElevationTransformation(planet, &transformation)

		for k := 0; k < 10; k++ {
			transformation := randomElevationTransformation(subdivisions / 2)
			applyElevationTransformation(planet, &transformation)
		}
	}
}
