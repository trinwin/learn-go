package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	// range will keep looping over the channel until it's closed
	for v := range c{
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // close a channel when done
	// closing the c channel on line 8
}