package main

// 	"errors"

// func main() {
// 	err:=customError()
// 	if err!=nil{
// 		wrappedErr := fmt.Errorf("context: %w", err)
// 		fmt.Println("wrapped error:", wrappedErr)

// 		fmt.Println("Original error:",errors.Unwrap(wrappedErr))
// 		originalError:=errors.Unwrap(wrappedErr)
// 		fmt.Println("Original error:",originalError)

// 		fmt.Println("Unwrap original error:",errors.Unwrap(wrappedErr))
// 	}else{
// 		fmt.Println("no error")
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

