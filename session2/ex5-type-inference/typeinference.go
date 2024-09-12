package main

import "fmt"

func main() {
	x := 10 //(Go infers x as int)
	fmt.Println(x)
	name := "Golang"
	fmt.Println(name)
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
	var y int      //(zero value is 0)
	fmt.Println(y)
	var z int = 10 //(forces x to be int)
	fmt.Println(z)
	fmt.Println(getNumber())
}

func getNumber() int { return 42 }