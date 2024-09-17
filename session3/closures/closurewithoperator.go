package main

// import "fmt"

func createOperator(op string) func(int, int) int {
	return func(a int, b int) int {
		switch op {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			return 0
		}
	}
}

// func main() {
// 	add := createOperator("+")
// 	del:=createOperator("-")
// 	fmt.Println(add(3, 4)) // Output: 7
// 	fmt.Println(del(4,3))
// }