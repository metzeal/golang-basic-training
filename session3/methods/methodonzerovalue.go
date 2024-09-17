package main

// import "fmt"

type Person struct {
	Name string
}

func (p Person) Greet() string {
	if p.Name == "" {
		return "Hello, stranger!"
	}
	return "Hello, " + p.Name + "!"
}
// func main() {
// 	var p Person           // Zero value of Person
// 	fmt.Println(p.Greet()) // Output: Hello, stranger!
// }
