package main

import (
	"fmt"
	"strings"
)

type IntProcessor func(int) int

func double(n int) int {
	return n * 2
}

func main() {
	var processor IntProcessor = double
	result := processor(5) // Output: 10
	fmt.Println("result:", result)

	var manipulator StringManipulator = toUpperCase
	strresult := manipulator("hello") // Output: HELLO
	fmt.Println("strresult:", strresult)

	var myFunc func(int) int // Zero value is nil

	if myFunc == nil {
		fmt.Println("Function is not initialized") // Output: Function is not initialized
	}

}

type StringManipulator func(string) string

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}
