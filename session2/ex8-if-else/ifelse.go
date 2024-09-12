package main

import "fmt"

func main() {
	number := -1
	if number > 0 {
		fmt.Println("number is positive")
	} else if number < 0 {
		fmt.Println("number is negative")
	} else {
		fmt.Println("number is zero")
	}
}