package main

import "fmt"

func main() {
// A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, 
// either because the surrounding function executed a return statement, reached the end of its function body, or 
// because the corresponding goroutine is panicking.
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
