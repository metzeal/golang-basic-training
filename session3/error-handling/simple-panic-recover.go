package main

import "fmt"

// func main() {

// 	example()
// }

func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}
