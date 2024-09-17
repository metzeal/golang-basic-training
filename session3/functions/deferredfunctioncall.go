package main

import "fmt"

// import "fmt"

// func main() {
// 	fmt.Println("Start")
// 	defer cleanup()
// 	fmt.Println("Exiting..")
// }

func cleanup() {
	defer fmt.Println("This will be printed last")
	fmt.Println("Cleanup resources..") 
}