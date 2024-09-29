package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {

// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	go func(ctx context.Context) {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Goroutine canceled")
// 		}
// 	}(ctx)
// 	time.Sleep(5 * time.Second)
// }