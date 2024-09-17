package main

// import "fmt"

func defaultGreeting() func(string) string {
	greeting := "Hello"
	return func(name string) string {
		if name == "" {
			return greeting + ", World!"
		}
		return greeting + ", " + name + "!"
	}
}

// func main() {
// 	greet := defaultGreeting()
// 	fmt.Println(greet(""))      // Output: Hello, World!
// 	fmt.Println(greet("Alice")) // Output: Hello, Alice!
// }
