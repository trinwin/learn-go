package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, playground", 42, true)

	fmt.Println(n)
	fmt.Println(err)

	// Or if we want to ignore err (cannot have unused variables in Go)
	nn, _ := fmt.Println("Hello, playground", 42, true)

	fmt.Println(nn)
}
