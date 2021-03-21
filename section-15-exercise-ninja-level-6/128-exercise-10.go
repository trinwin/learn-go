package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	
}

// create a func which “encloses” the scope of a variable
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}