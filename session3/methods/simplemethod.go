package main

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// func main() {
// 	rect := Rectangle{10, 20}
// 	fmt.Println("Area of Rectangle:",rect.Area())
// }