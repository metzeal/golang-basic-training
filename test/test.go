// package main

// import "fmt"

// func main() {

// 	mychan := make(chan string)

// 	go sender(mychan, "message")

// 	// s := <-mychan
// 	// fmt.Println(s)	
// 	receiver(mychan)
// }

// func receiver(mychan chan string) {
// 	s := <-mychan
// 	fmt.Println(s)	
// }
// func sender(mychan chan string, message string) {
// 	mychan <- message
// }