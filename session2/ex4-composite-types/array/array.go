package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30
	fmt.Println("array contents:",array)
	fmt.Println("array length: ",len(array))

	var copyarray = array
	if copyarray== array{
		fmt.Println("arrays are same")
	}else{
		fmt.Println("arrays are not same")
	}


}