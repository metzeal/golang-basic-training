package main

import "fmt"

func identify(i interface{}) {
	switch v := i.(type) {
	case float64:
		fmt.Printf("Float64: %f\n", v)
	case bool:
		fmt.Printf("Bool: %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

// func main() {
// 	identify(3.14) // Output: Float64: 3.140000
// 	identify(true) // Output: Bool: true
// 	identify("hello")
// }
