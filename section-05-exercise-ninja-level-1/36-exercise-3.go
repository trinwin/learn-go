package main

import "fmt"

// Declare and assign value to var variable at a package level
var x int = 42
var y string = "James Bond"
var z bool  = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z )
	fmt.Println(s)
}
