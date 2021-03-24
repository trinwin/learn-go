package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incre := 0

	const gr = 10

	var wg sync.WaitGroup
	wg.Add(gr)

	// Create a mutex
	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func ()  {
			mu.Lock()
			temp := incre
			temp++
			incre = temp
			mu.Unlock()
			wg.Done()
			
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Increment:", incre)
}

