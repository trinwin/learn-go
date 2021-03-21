package main

import "fmt"

func main()  {
	a := 42
	fmt.Println(a)  // value 
	fmt.Println(&a)	// address
	fmt.Printf("%T\n", a) 	// int
	fmt.Printf("%T\n", &a) 	// pointer to an int - *int

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * is an operator
	fmt.Println(*&a) // * is an operator
	// * with a type - pointer
	// * with an identifier - dereference a address to get the value

	*b = 43 // == a
	fmt.Println(b) // same address - different value now
	fmt.Println(*b)
}