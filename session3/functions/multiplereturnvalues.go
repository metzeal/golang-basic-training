package main

import "fmt"

// func main() {
// 	result, err := multipleReturnValues(10, 2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }
func multipleReturnValues(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}
