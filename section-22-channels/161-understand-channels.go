package main

import "fmt"

func main()  {
	c := make(chan int) // Allocate a channel

	go func ()  {
		c <- 42 // putting values on a channel
	}()
	
	
	fmt.Println(<-c) // Taking values off of a channel
	
}

/*
	c := make(chan int) // Allocate a channel

	// Send and Receive have to happen at the same time
	// Channel block example
	c <- 42 // Send a signal; Value does not matter
	
	fmt.Println(<-c) // Taking values off of a channel
*/