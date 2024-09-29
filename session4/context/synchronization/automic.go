package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 // Shared variable (int64 is required for atomic operations)

func increment(wg *sync.WaitGroup) {
	// Atomically increment the counter
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Start 5 goroutines that increment the counter concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Output the final value of the counter
	fmt.Printf("Final Counter Value: %d\n", counter)
}
