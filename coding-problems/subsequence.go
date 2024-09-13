package main

import "fmt"

func solution(a string, b string) int{
	i, j := 0, 0
	count := 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			count++
			i++
			j++
		} else {
			if len(a) > len(b) {
				i++
			} else {
				j++
			}
		}
	}
	return count
}
func main() {
	result:=solution("abcdef","ace")
	fmt.Print(result)

	// result=solution("ace","ace")
	// fmt.Print(result)

	// result=solution("ace","ad")
	// fmt.Print(result)

	// result=solution("ace","pqr")
	// fmt.Print(result)

	result=solution("abaaba","babbab")
	fmt.Print(result)

}