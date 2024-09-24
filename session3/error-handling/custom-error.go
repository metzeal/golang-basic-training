package main

// import "fmt"

// func main() {
// 	err:=customError()
// 	if err!=nil{
// 		fmt.Println("Error:", err)
// 	}else{
// 		fmt.Println("No Error")
// 	}
// }

// type MyError struct {
// 	Code    int
// 	Message string
// }

// func customError() error{
// 	return &MyError{Code: 123, Message: "custom error"}
// }
// func (e *MyError) Error() string {
// 	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
// }
