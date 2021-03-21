package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}
	sum := foo(a...)
	fmt.Println(sum)
	
	sum2 := bar(a)
	fmt.Println(sum2)
}

func foo(a ...int) int {
	sum := 0	
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar(b []int) int {
	sum := 0	
	for _, v := range b {
		sum += v
	}
	return sum
}