package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var incre int64

	const gr = 10

	var wg sync.WaitGroup
	wg.Add(gr)


	for i := 0; i < gr; i++ {
		go func ()  {
			atomic.AddInt64(&incre, 1)
			runtime.Gosched()
			fmt.Println("Increment\t", atomic.LoadInt64(&incre))
			wg.Done()
			
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Increment:", incre)
}

