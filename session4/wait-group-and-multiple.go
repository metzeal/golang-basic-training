package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("I am done, Thanks for waiting")
		wg.Done()
	}()
	go func() {
		fmt.Println("I am also done, Thanks for waiting")
		wg.Done()
	}()	
	wg.Wait()
	fmt.Println("I am exiting now")
}
