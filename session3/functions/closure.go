package main

// func main() {
// 	increment := counter()
// 	fmt.Println(increment()) // Output: 1
// 	fmt.Println(increment()) // Output: 2
// }

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
