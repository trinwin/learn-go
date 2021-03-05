package main

import "fmt"

// Assign a type to a newly created type
type kem int
var x kem

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
