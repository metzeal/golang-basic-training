package main

import "fmt"

func main() {

	basicSwitch()
	expressionLessSwitch()

	fallthroughSwitch()
	
	multipleCasesSwitch()
	
	typeSwitch()
	

}

type MyInterface interface{
	ToString() string

}

type Student struct{
	Name string
	Age int	
}

func (s Student) ToString() string{
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

func typeSwitch() {

	student := Student{Name: "Rama", Age: 10}
	var i MyInterface=student
	switch v := i.(type) {
	case Student:
		fmt.Println("Student", v)
	default:
		fmt.Println("Unknown type")
	}
}

func multipleCasesSwitch() {
	x := 2
	switch x {
	case 1, 2, 3:
		fmt.Println("One, Two, or Three")
	}
}

func fallthroughSwitch() {
	x := 1
	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
}

func expressionLessSwitch() {
	x := -1
	switch {
	case x > 0:
		fmt.Println("Positive")
	case x < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}

func basicSwitch() {
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}
}