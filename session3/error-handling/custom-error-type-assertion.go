package main

import "fmt"


func main() {
	err:=customError()
	handleError(err)
}

type MyError struct {
	Code    int
	Message string
}

func customError() error{
	return &MyError{Code: 123, Message: "custom error"}
}
func (e *MyError) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}


func handleError(err error) {
    if myErr, ok := err.(*MyError); ok {
        fmt.Printf("Custom error occurred: %s\n", myErr.Error())
        // Handle specific error based on Code or Message
    } else {
        fmt.Println("General error:", err)
    }
}
