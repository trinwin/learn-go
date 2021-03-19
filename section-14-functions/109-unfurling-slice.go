package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// x := sum(xi)
	x := calsum(xi...) //unfurling a slice of int
	// in this case no new slice is created, the same underlying array will be used
	fmt.Println("The total is", x)
}

// variadic parameter ... 
// ... parameter has to be the last one
func calsum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice of int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
