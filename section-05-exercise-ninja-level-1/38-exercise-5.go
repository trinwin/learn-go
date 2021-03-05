package main

import "fmt"

// Assign a type to a newly created type
// int is the underlying type of custom type kem
type kem int
var x kem
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// Convert type of x to type of y (int) before assign
	// Without conversion we have error "cannot use x (type kem) as type int in assignment"
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
