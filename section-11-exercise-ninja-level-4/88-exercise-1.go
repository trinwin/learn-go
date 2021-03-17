package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of a is %T", a)
}
