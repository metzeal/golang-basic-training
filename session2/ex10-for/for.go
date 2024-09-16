package main

import "fmt"

func main() {
	indexBased()
}

func indexBased() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}

func conditionOnly() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func infite() {
	for {
		fmt.Println("Running")
		break
	}

}

func rangeBased() {
	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Println(index, value)
	}

}

func forContinue() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func forBreak() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	
}