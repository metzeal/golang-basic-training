package main

import (
	"fmt"
	"strconv"
)


func main() {
	primitiveconversion()
	stringconversion()
	customConversion()
}

func customConversion() {
	type MyInt int
	var i MyInt = MyInt(42)
	fmt.Println("i:", i)

	var j int = int(i)
	fmt.Println("j:", j)

}
func stringconversion() {
	i, _ := strconv.Atoi("42")
	fmt.Println("i:", i)
	s := strconv.Itoa(42)
	fmt.Println("s:", s)
}

func primitiveconversion() {
	var x int = 42
	fmt.Println("x:", x)
	var y float64 = float64(x)
	fmt.Print("y:", y)
	i := int(y)
	fmt.Print("i:", i)
}

