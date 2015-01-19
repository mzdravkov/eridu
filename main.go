package main

// import "fmt"

func main() {
	triLines := trianglesInLines(5)
	planet := generateNewRegions(triLines)
	transformPlanetRelief(&planet)

	write(&planet)

	// for i := 0; i < 1536; i++ {
	// 	o := len(*(planet[i]))
	// 	// fmt.Println((*planet[i])[o-1])
	// 	for k := int16(0); k < (*planet[i])[o-1].Elevation; k++ {
	// 		fmt.Print(".")
	// 	}
	// 	fmt.Println("")
	// }
}
