package main

import "testing"

func TestTrianglesInLines(t *testing.T) {
	result := trianglesInLines(3)
	expected := []int{5, 15, 25, 35, 40, 40, 40, 40, 35, 25, 15, 5}

	if len(result) != len(expected) {
		t.Error("Test is failing for trianglesInLines with subdivisions = 3")
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Error("Test is failing for trianglesInLines with subdivisions = 3")
		}
	}
}

func TestSubdivisionsFromLines(t *testing.T) {
	if subdivisionsFromLines(trianglesInLines(3)) != 3 {
		t.Error("subdivisionsFromLines doesn't return the right result.")
	}
}

func TestGenerateNewRegions(t *testing.T) {
	layers := generateNewRegions(trianglesInLines(3))

	if len(*layers[0]) != 20 {
		t.Error("generateNewRegions doesn't work.")
	}

	if len(*layers[1]) != 80 {
		t.Error("generateNewRegions doesn't work.")
	}

	if len(*layers[2]) != 320 {
		t.Error("generateNewRegions doesn't work.")
	}

	newRegion := Region{0, 0}
	for i := 0; i < len(*layers[0]); i++ {
		if (*layers[0])[i] != newRegion {
			t.Error("generateNewRegions doesn't work.")
		}
	}
}
