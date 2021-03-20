package main

import "fmt"

func main() {
	// anonymous function - self executing function
	func ()  {
		fmt.Println("anonymous func ran")
	}()
	func (x int)  {
		fmt.Println("number:", x)
	}(42)
}
