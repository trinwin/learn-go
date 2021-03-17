package main

import (
	"fmt"
)

func main() {
	r1 := []string{"James", "Bond", "Shaken, not stirred"}
	r2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(r1)
	fmt.Println(r2)

	table := [][]string{r1, r2}
	fmt.Println(table)

	for i, row := range table {
		fmt.Println("record: ", i)
		for j, val := range row {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
