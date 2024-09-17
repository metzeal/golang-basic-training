package main

// import "fmt"

// func main() {
// 	resizableRectangle := ResizableRectangle{10.2, 22.9}
// 	fmt.Println("Rectangle:",resizableRectangle)
// 	resizableRectangle.resize(12, 13)
// 	fmt.Println("Resized Rectangle:",resizableRectangle)

// }
type ResizableRectangle struct {
	Width, Height float64
}

func (r *ResizableRectangle) resize(newWidth, newHeight float64) {
	r.Width = newWidth
	r.Height = newHeight

}
