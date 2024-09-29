package main

// var (
// 	buffer    []int      // Shared resource (buffer)
// 	maxSize   = 5        // Max buffer size
// 	mutex     sync.Mutex // Mutex to protect the buffer
// 	cond      = sync.NewCond(&mutex) // Condition variable
// )

// func produce(id int, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second) // Simulate production time

// 		mutex.Lock()
// 		for len(buffer) == maxSize {
// 			fmt.Printf("Producer %d: Buffer full, waiting...\n", id)
// 			cond.Wait() // Wait if the buffer is full
// 		}

// 		// Produce an item and add it to the buffer
// 		buffer = append(buffer, i)
// 		fmt.Printf("Producer %d: Produced %d, buffer: %v\n", id, i, buffer)

// 		cond.Signal() // Signal to consumer
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }

// func consume(id int, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		mutex.Lock()
// 		for len(buffer) == 0 {
// 			fmt.Printf("Consumer %d: Buffer empty, waiting...\n", id)
// 			cond.Wait() // Wait if the buffer is empty
// 		}

// 		// Consume an item from the buffer
// 		item := buffer[0]
// 		buffer = buffer[1:]
// 		fmt.Printf("Consumer %d: Consumed %d, buffer: %v\n", id, item, buffer)

// 		cond.Signal() // Signal to producer
// 		mutex.Unlock()

// 		time.Sleep(time.Second) // Simulate consumption time
// 	}
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Start 1 producer and 1 consumer
// 	wg.Add(1)
// 	go produce(1, &wg)

// 	wg.Add(1)
// 	go consume(1, &wg)

// 	wg.Wait()
// 	fmt.Println("All tasks completed.")
// }
