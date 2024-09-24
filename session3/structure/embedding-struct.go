package main

import "fmt"

type Author struct {
	Name  string
	Birth int
}
type Book struct {
	Author // Embedding Author struct
	Title  string
	Year   int
}
func (a Author) Bio() string {
    return fmt.Sprintf("%s (born %d)", a.Name, a.Birth)
}

func main() {
	b := Book{Author: Author{Name: "George Orwell", Birth: 1903}, Title: "1984", Year: 1949}
	fmt.Println(b.Name) 
	
	fmt.Println(b.Bio())
}