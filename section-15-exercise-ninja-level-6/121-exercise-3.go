package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("function #1")
}

func bar() {
	fmt.Println("function #2")
}