package main

import "fmt"

type Counter struct {
	count int
}

// Method with value receiver
func (c Counter) Increment() {
	c.count++
}

func main() {
	c := Counter{count: 0}
	c.Increment()
	fmt.Println("Count:",c.count) // Output: 0
}
