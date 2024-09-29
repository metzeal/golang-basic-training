package main

// var (
// 	counter int       // Shared resource
// 	mutex   sync.Mutex // Mutex to protect the shared resource
// )

// func increment(wg *sync.WaitGroup) {
// 	// Lock the mutex before accessing the shared resource
// 	mutex.Lock()
// 	counter++
// 	fmt.Println("Counter:", counter)
// 	// Unlock the mutex after accessing the shared resource
// 	mutex.Unlock()
// 	wg.Done() // Indicate the goroutine is finished
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go increment(&wg) // Launching multiple goroutines
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()
// 	fmt.Println("Final Counter Value:", counter)
// }
