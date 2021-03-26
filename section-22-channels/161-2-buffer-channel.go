package main

import "fmt"

func main()  {
	// Buffer channel allows value to sit in there
	c := make(chan int, 2) // Allocate a channel

	// 42 is sitting on the 1 buffer channel we declare
	c <- 42 // putting values on a channel
	// c <- 43 // this won't work since we only have 1 buffer channel for 42
					// to make it work, we have to change num of buffers to 2

	fmt.Println(<-c) // Taking values off of a channel
	// fmt.Println(<-c) // Taking values off of a channel
	
}
