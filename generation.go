package main

import (
	"math"
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

	triangles = 10*segmentSize
	for i := segmentSize; i < 2*segmentSize; i++ {
		result[i] = triangles
	}

	triangles = 10*segmentSize - 5
	for i := 2*segmentSize; i < 3*segmentSize; i++ {
		result[i] = triangles
		triangles -= 10
	}

	return result
}

struct Region {
	Type uint8
	Elevation int16
}

func generateNewRegions(lines []int) error {
	
}