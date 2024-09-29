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

// 	Loop:
// 	for {
// 		select {
// 		case msg, ok := <-ch:
// 			if ok {
// 				fmt.Println(msg)
// 			} else {
// 				fmt.Println("Channel closed")
// 				break Loop
// 			}
// 		default:
// 			fmt.Println("No message")
// 		}
// 	}
// }
