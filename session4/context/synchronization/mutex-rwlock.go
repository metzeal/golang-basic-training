package main

// var (
// 	counter int          // Shared resource
// 	rwMutex sync.RWMutex // RWMutex to manage read/write access
// )

// // Function to read the value of counter
// func readCounter(id int, wg *sync.WaitGroup) {
// 	rwMutex.RLock() // Acquire read lock
// 	fmt.Printf("Goroutine %d reading: %d\n", id, counter)
// 	time.Sleep(1 * time.Second) // Simulate some reading time
// 	rwMutex.RUnlock()           // Release read lock
// 	wg.Done()                   // Indicate the goroutine is finished
// }

// // Function to write to the counter
// func writeCounter(wg *sync.WaitGroup) {
// 	rwMutex.Lock() // Acquire write lock
// 	counter++
// 	fmt.Println("Writing new value:", counter)
// 	time.Sleep(1 * time.Second) // Simulate some writing time
// 	rwMutex.Unlock()            // Release write lock
// 	wg.Done()                   // Indicate the goroutine is finished
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Start 5 reader goroutines
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go readCounter(i, &wg)
// 	}

// 	// Start 1 writer goroutine
// 	wg.Add(1)
// 	go writeCounter(&wg)

// 	// Wait for all goroutines to finish
// 	wg.Wait()
// 	fmt.Println("Final Counter Value:", counter)
// }
