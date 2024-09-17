package main

import "fmt"

func createMap() (func(string, int), func(string) (int, bool), func(string, int), func(string)) {
	m := make(map[string]int)

	add := func(key string, value int) {
		m[key] = value
	}
	get := func(key string) (int, bool) {
		value, ok := m[key]
		return value, ok
	}
	update := func(key string, value int) {
		if _, exists := m[key]; exists {
			m[key] = value
		}
	}
	delete := func(key string) {
		delete(m, key)
	}
	return add, get, update, delete
}

func main() {
	add, get, update, delete := createMap()
	add("key1", 10)
	val, _ := get("key1")
	fmt.Println(val) // Output: 10
	update("key1", 20)
	val, _ = get("key1")
	fmt.Println(val) // Output: 20
	delete("key1")
	_, ok := get("key1")
	fmt.Println(ok) // Output: false
}