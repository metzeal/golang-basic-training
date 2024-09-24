package main

import "fmt"

// func main() {
// 	fmt.Println("Start")
// 	safeFunction()
// 	fmt.Println("End")
// }
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	fmt.Println("Inside safe function")
	riskyOperation()
}

func riskyOperation() {
	panic("Critical error!")
}
