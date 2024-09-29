package main

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 			fmt.Println("Sent:", i)
// 		}
// 		close(ch)
// 	}()

// 	for{
// 		select {
// 		case msg := <-ch:
// 			fmt.Println(msg)
// 		default:
// 			fmt.Println("No message")
// 		}
// 	}

// }
