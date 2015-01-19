package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func write(planet *[]*[]Region) {
	csvfile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	for i := 0; i < len(*planet); i++ {
		row := make([]string, len(*(*planet)[i]))
		for k := 0; k < len(*(*planet)[i]); k++ {
			row[k] = strconv.Itoa(int((*(*planet)[i])[k].Elevation))
		}
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	}
	writer.Flush()
}
