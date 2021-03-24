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

	for i := 0; i < gr; i++ {
		go func ()  {
			temp := incre
			runtime.Gosched()
			temp++
			incre = temp
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Increment:", incre)
}

