package main

import "fmt"

// declare + assign = initialization
var y = 43

// Scope of y is from line 5 to 18 while x is only within function main

// Decalre there is a variable with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPe int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for boolean, 0 for int, 0.0 for float, "" for string
// nil for pointer, func, interfaces, slies, channels and maps.
var z int

func main() {
	// Short declaration operator
	x := 42 // Declare and assign a value to a variable
	fmt.Println(x)
	// Var keyword

	fmt.Println(y)
	foo()
	// You can only declare a variable outside of functions body using VAR
	// Best practice: Keep the scope as "narrow" as possible

	fmt.Println(z)
}

func foo() {
	fmt.Println("again y:", y)

}
