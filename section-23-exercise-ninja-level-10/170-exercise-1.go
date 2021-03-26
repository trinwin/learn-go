package main

import "fmt"

func main()  {
	c := make(chan int) // Allocate a channel

	go func ()  {
		c <- 42 // putting values on a channel
	}()
	
	
	fmt.Println(<-c) // Taking values off of a channel
	
}