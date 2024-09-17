package main

// func main() {
// 	add := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println("add(10,20):", add(10, 20))
// }

func applyOperatoin(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}
