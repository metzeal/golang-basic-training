package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	printNumbers(10, 20, 20, 30)
// 	result := joinStrings(", ", "apple", "banana", "cherry")
// 	fmt.Println(result) // Output: apple, banana, cherry
// 	summarize("Numbers:", 1, 2, 3, 4, 5)
// 	printValues("Hello", 42, 3.14, true)
// }

func printNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}


func joinStrings(separator string, values ...string) string {
    return strings.Join(values, separator)
}

func summarize(prefix string, values ...int) {
    fmt.Println(prefix)
    for _, value := range values {
        fmt.Println(value)
    }
}


func printValues(values ...interface{}) {
    for _, value := range values {
        fmt.Println(value)
    }
}

