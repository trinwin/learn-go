package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	
	a = append(a, 52)
	fmt.Println(a)
	
	a = append(a, 53, 54, 55)
	fmt.Println(a)

	b := []int{56, 57, 58, 59, 60}
	a = append(a, b...)
	fmt.Println(a)	
}
