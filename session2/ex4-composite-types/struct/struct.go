package main

import "fmt"

type Student struct {
	Name string;	Age  int
}

func main() {
	student := Student{Name: "Rama", Age: 10}
	fmt.Println("student:", student)
	fmt.Println("student name:", student.Name)
	fmt.Println("student age:", student.Age)
	teacher:=new(Student)
	fmt.Println("teacher:", *teacher)

	students:=map[string]Student{
		"Rama": {Name: "Rama", Age: 10},
		"Shyam":{Name: "Shyam", Age: 20},
		"Gopi": {Name: "Gopi", Age: 30},
	}
	fmt.Println("students:", students)
}

func (s *Student) getAge() int {
	return s.Age
}


func (s *Student) setAge(age int) {
	s.Age = age
}

func (s *Student) getName() string {
	return s.Name
}

func (s *Student) setName(name string) {
	s.Name = name
}