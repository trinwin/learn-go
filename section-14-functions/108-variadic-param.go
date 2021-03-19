package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5)
	fmt.Println("The total is", x)
}

// variadic parameter ...
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice of int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
