package main

import "fmt"

// func main() {
// 	result, err := noError()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}else{
// 		fmt.Println("Result:", result)
// 	}

// 	result, err=returnError()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}else{
// 		fmt.Println("Result:", result)
// 	}
// }

func noError() (int, error) {
	return 0, nil
}


func returnError() (int, error) {
	return 0, fmt.Errorf("error")
}
