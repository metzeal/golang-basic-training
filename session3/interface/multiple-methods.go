package main

import "fmt"
type Speaker interface {
    Speak() string
}
type Animal struct {
    Name string
}
func (a Animal) Speak() string {
    return "Animal noise"
}

type Dog struct {
    Animal
}

func (d Dog) Speak() string {
    return "Woof"
}

func main() {
    var s Speaker
	animal:=Animal{Name: "Rex"}
    s = Dog{animal}

    fmt.Println(s.Speak())  // Output: Woof
	fmt.Println(animal.Speak())
}
