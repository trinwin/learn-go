package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for k := 1; k <= 3; k++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Println("")
	}
}
