package main

import "fmt"

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	go sender(ch)
// 	receiver(ch)
// }

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}