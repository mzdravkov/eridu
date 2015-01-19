package main

import "fmt"

func main() {
	triLines := trianglesInLines(4)
	planet := generateNewRegions(triLines)
	transformPlanetRelief(&planet)

	// write(&planet)

	for i := 0; i < 1536; i++ {
		fmt.Println(planet[i])
	}
}
