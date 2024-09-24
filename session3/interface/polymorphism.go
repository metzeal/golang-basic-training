package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shaper) {
	fmt.Println("Area:", s.Area())
}
// func main() {
// 	c := Circle{Radius: 5}
// 	r := Rectangle{Width: 4, Height: 6}

// 	printArea(c) // Output: Area: 78.5
// 	printArea(r) // Output: Area: 24
// }
