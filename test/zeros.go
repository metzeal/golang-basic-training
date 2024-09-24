package main

import "fmt"

func main() {
	input := [8]int{1, 2, 0, 4, 3, 0, 5, 0}
	for i:=0; i<len(input)-1;i++ {
		for j:=0;j<len(input)-1;j++{
			if input[j] == 0 {
				input[j] = input[j+1]
				input[j+1] = 0
			}	
		}
	}
	fmt.Print(input)
}