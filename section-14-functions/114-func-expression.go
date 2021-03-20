package main

import "fmt"

// Assign a func to a variable
func main() {
	f := func ()  {
		fmt.Println("my func expression")
	}

	f() // call the function

	g := func (x int)  {
		fmt.Println("my birth month:", x)
	}

	g(07)
}
