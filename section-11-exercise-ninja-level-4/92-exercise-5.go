package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	
	a = append(a[:3], a[6:]...)
	fmt.Println(a)
}
