package main

import "fmt"

func main() {
	foo(callback)
}

func callback()  {
	fmt.Println("this is a callback")
}

func foo(f func()) {
	f()
}

