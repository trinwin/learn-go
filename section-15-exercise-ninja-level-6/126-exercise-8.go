package main

import "fmt"

func main() {
	x := foo()
	x()
}

func foo() func() int  {
	return func() int{
		fmt.Println("Calvin is stupid")
		return 0
	}
}

