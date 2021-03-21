package main

import "fmt"

func main()  {
	x := 0
	fmt.Println(&x)
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)
	fmt.Println(&x)
}

func foo(y *int)  {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}