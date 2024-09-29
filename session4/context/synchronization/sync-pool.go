package main

// // Define a sync.Pool for object reuse
// var pool = sync.Pool{
// 	New: func() interface{} {
// 		fmt.Println("Allocating new object...")
// 		return new(int) // The object we want to pool (in this case, an integer)
// 	},
// }

// func worker(id int, wg *sync.WaitGroup) {
// 	// Get an object from the pool
// 	obj := pool.Get().(*int)
// 	*obj = id // Use the object

// 	fmt.Printf("Worker %d using object with value: %d\n", id, *obj)

// 	time.Sleep(time.Second) // Simulate some work

// 	// Put the object back into the pool for reuse
// 	pool.Put(obj)

// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Create 5 worker goroutines that use objects from the pool
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers are done.")
// }
