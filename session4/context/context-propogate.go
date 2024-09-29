package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	go worker(ctx)
// 	time.Sleep(5 * time.Second)
// }

// func worker(ctx context.Context) {
//     for {
//         select {
//         case <-ctx.Done():
//             return
//         default:
//             fmt.Print("Working...")
//         }
//     }
// }
