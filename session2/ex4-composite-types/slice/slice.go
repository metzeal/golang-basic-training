package main

import "fmt"

func main() {
	var numbers []int = []int{1, 2, 3, 4}

	squares := []int{numbers[0] * numbers[0], numbers[1] * numbers[1], numbers[2] * numbers[2], numbers[3] * numbers[3]}

	fmt.Println("numbers:",numbers)
	fmt.Println("squares:",squares)

	cubes:=make([]int,0)

	cubes = append(cubes, numbers[0] * numbers[0] * numbers[0])
	fmt.Println("cubes:",cubes)
	fmt.Println("cubes length:",len(cubes))
	fmt.Println("cubes capacity:",cap(cubes))

	cubes = append(cubes, numbers[1] * numbers[1] * numbers[1])
	fmt.Println("cubes:",cubes)
	fmt.Println("cubes length:",len(cubes))
	fmt.Println("cubes capacity:",cap(cubes))
	// capacity is doubled new array is allocated
	cubes = append(cubes, numbers[2] * numbers[2] * numbers[2])
	fmt.Println("cubes:",cubes)	
	fmt.Println("cubes length:",len(cubes))
	fmt.Println("cubes capacity:",cap(cubes))

	cubes = append(cubes, numbers[3] * numbers[3] * numbers[3])
	fmt.Println("cubes:",cubes)	
	fmt.Println("cubes length:",len(cubes))
	fmt.Println("cubes capacity:",cap(cubes))

	fmt.Println("first half cubes:",cubes[0:2])

	fmt.Println("remaining half cubes:",cubes[2:4])

}