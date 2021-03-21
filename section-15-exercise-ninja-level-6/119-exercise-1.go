package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x,y,z)
}

func foo() int  {
	return 1
}

// 2 or more return values then use (a, b) format
func bar() (int, string) {
	return 1, "hello"
}
