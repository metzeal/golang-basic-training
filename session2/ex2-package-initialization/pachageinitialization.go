package main

import (
	_ "ex2-package-initialization/configuration"
	_ "ex2-package-initialization/database"
	"fmt"
)

/*
main functis entry point
*/


func main() {
	if true{
		fmt.Println("hello")
	}
	fmt.Println("main called")
	
}