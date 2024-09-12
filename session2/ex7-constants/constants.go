package main

import "fmt"

const pi = 3.14
const name = "PI"

const (
	RED = iota
	BLUE
	YELLOW
)

const (
	MON = 1
	TUE
	WED
)

func main() {
	fmt.Println(pi)
	fmt.Println(name)
	fmt.Println(RED)
	fmt.Println(YELLOW)
	fmt.Print(TUE)
}