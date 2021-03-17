package main

import "fmt"

func main()  {
	// arrays are building block for slides
	// x := type{}  // composite literal
	x := []int{4,5,6}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range x {
		fmt.Println(i,v)
	}

	// Slicing a slice
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// Append to a slice
	 x = append(x, 77, 88, 99, 1014)
	 fmt.Println(x)

	 y := []int{234, 456, 678, 987}
	 x = append (x, y...)
	 fmt.Println(x)

	//  Deleting from a slice
	//  Use append to delete from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// Make - Create slice with a specified size
	// if slice grows, a new array is created for a new slice and the old one is thrown away
	// which takes processing power so use make to create a slice with a designated size 
	// if you know the size beforehand
	a := make([]int, 10, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a[0] = 42
	a[9] = 99
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	// a[10] = 10 - this doesn't work because we declare a slice with len 10 (0-9) 
	x = append(a, 3423)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Multi-dimensional slice
	n := []string{"a", "b", "c", "d"}
	fmt.Println(n)
	m := []string{"x", "y", "z", "-"}
	fmt.Println(m)

	nm := [][]string{n, m}
	fmt.Println(nm)
}
