package main

import "fmt"

// return a func from a func
func main() {

	// x := bar()
	// fmt.Println(x)
	// fmt.Printf("%T\n", x)

	// i := x()
	// fmt.Println(i)
	// OR
	// fmt.Println(x())
	// OR
	fmt.Println(bar()())
}


func bar() func() int  {
	return func() int {
		return 451
	}
}