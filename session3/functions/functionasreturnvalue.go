package main

// func main() {
// 	func() {
// 		fmt.Println("This is an anonymous function executed immediately")
// 	}()

// 	multiplier := makeMultiplier(3)
// 	fmt.Println(multiplier(4)) // Output: 12
// }

func makeMultiplier(factor int) func(int) int {
	
	return func(x int) int {
		return x * factor
	}
}
