package main

import "fmt"
type Person struct {
    Name    string
    Age     int
}
func (p *Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}
// func main() {
//     p := Person{Name: "Bob", Age: 25}
//     fmt.Println(p.Greet())
// }
