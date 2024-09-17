package main

// func main() {
// 	totalPrice := calculateTotalPrice(100.0, 0.2)
// 	fmt.Println("totalPrice", totalPrice)
// }

func calculateTotalPrice(price float64, taxRate float64) float64 {
	return price + (price * taxRate)
}

