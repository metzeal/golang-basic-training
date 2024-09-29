package main

// var (
// 	once       sync.Once
// 	connection *DBConnection
// )

// // DBConnection simulates a database connection
// type DBConnection struct {
// 	connected bool
// }

// // Simulate initializing a database connection (should only happen once)
// func initializeDBConnection() *DBConnection {
// 	fmt.Println("Initializing database connection...")
// 	time.Sleep(2 * time.Second) // Simulate time to establish a connection
// 	return &DBConnection{connected: true}
// }

// // GetDBConnection returns the initialized DB connection, using sync.Once to ensure it's initialized once
// func GetDBConnection() *DBConnection {
// 	once.Do(func() {
// 		connection = initializeDBConnection()
// 	})
// 	return connection
// }

// // Simulate a worker trying to use the database connection
// func worker(id int, wg *sync.WaitGroup) {
// 	conn := GetDBConnection() // Try to get the DB connection
// 	fmt.Printf("Worker %d using database connection: %v\n", id, conn.connected)
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Launch multiple goroutines that try to access the database connection
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers have finished using the database connection.")
// }
