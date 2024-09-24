package main

import (
	"errors"
	// "fmt"
)

// func main() {
// 	result, err:=returnNewError()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}else{
// 		fmt.Println("Result:", result)
// 	}
// }


func returnNewError() (int, error) {
	return 0, errors.New("new error")
}
