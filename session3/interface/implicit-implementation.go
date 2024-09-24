package main

type Runner interface {
	Run() string
}

type Athlete struct {
	Name string
}

func (a Athlete) Run() string {
	return a.Name + " is running"
}
// func main() {
// 	var r Runner = Athlete{Name: "Usain"}
// 	fmt.Println(r.Run()) // Output: Usain is running
// }
