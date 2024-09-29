package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func(ctx context.Context) {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Goroutine canceled")
// 		}
// 	}(ctx)
// 	time.Sleep(5 * time.Second)
// 	cancel() // Cancels the goroutine
// 	time.Sleep(15 * time.Second)

// }