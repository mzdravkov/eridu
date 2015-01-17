package main

import "testing"

func TestTrianglesInLines(t *testing.T) {
	result := trianglesInLines(3)
	expected := []int {5, 15, 25, 35, 40, 40, 40, 40, 35, 25, 15, 5} 

	if len(result) != len(expected) {
		t.Error("Test is failing for trianglesInLines with subdivisions = 3")
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Error("Test is failing for trianglesInLines with subdivisions = 3")
		}
	}
}