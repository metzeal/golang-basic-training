package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Create a buffered channel with a capacity of 3
// 	bufferedChan := make(chan int, 2)

// 	// Start a goroutine to receive data from the channel
// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			val := <-bufferedChan
// 			fmt.Println("Received:", val)
// 		}
// 	}()

// 	// Send data to the buffered channel
// 	bufferedChan <- 10
// 	fmt.Println("Sent 10")
// 	bufferedChan <- 20
// 	fmt.Println("Sent 20")
// 	bufferedChan <- 30
// 	fmt.Println("Sent 30")
// 	bufferedChan <- 40
// 	fmt.Println("Sent 40")
// 	// Sleep to allow goroutine to finish receiving
// 	time.Sleep(1 * time.Second)
// }
