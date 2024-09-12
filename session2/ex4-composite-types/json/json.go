package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {
	student := Student{Name: "Rama", Age: 10}

	student_bytes,err:=json.Marshal(student)
	if err!=nil{
		fmt.Println("Error:",fmt.Errorf("Error: %+v", err))
	}
	fmt.Println(student_bytes)
	fmt.Println(string(student_bytes))
	student_back:=Student{}
	json.Unmarshal(student_bytes, &student_back)
	fmt.Println(fmt.Sprintf("student_back: %+v", student_back))

	var studentmap map[string]interface{}
	json.Unmarshal(student_bytes, &studentmap)
	fmt.Println("studentmap:", studentmap)

	fmt.Print(studentmap["name"])
	fmt.Print(studentmap["Name"])
}