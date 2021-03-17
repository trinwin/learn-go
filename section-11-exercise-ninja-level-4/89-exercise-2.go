package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of a is %T", a)
}
