package main

import "fmt"

func main() {
	// var marks map[string]int = map[string]int{"Rama": 10, "Shyam": 20, "Gopi": 30}

	// marks:=make(map[string]int)
	var marks map[string]int=nil
	fmt.Println("marks:", marks)
	fmt.Println("marks length:", len(marks))
	if _, ok:=marks["raja"];ok{
		fmt.Println("marks:", marks["raja"])
	}else{
		fmt.Println("key not found")
		marks["raja"]=100
	}
	fmt.Println("marks:", marks)
	fmt.Println("marks length:", len(marks))

	delete(marks,"")
	fmt.Println("post delete non existing")

	fmt.Println("marks:", marks)
	fmt.Println("marks length:", len(marks))

	delete(marks,"Rama")
	fmt.Println("post delete existing")

	fmt.Println("marks:", marks)
	fmt.Println("marks length:", len(marks))
}