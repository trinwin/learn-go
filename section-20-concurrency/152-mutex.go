package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a race condition
func main()  {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func ()  {
			v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
}