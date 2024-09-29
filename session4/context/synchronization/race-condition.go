package main

// var counter = 0 // Shared variable (race condition)

// func increment(id int, wg *sync.WaitGroup) {
// 	for i := 0; i < 1000; i++ {
// 		// Incrementing the shared counter (no protection, race condition)
// 		counter++
// 	}
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Start 5 goroutines that increment the counter concurrently
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go increment(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final Counter Value: %d\n", counter)
// }
