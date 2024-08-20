package main

import "fmt"

func log(message string) {
	fmt.Println(message)
}

func add(a int, b int) (int, bool) {
	return a + b, true
}

func power(name string) (int, bool) {
	result := false
	if name == "Dan" {
		result = true
	}
	return 99, result
}

func main() {
	result, attempt := add(1, 1)
	name := "Dan"
	log(name)
	log(string(result))
	if attempt {
		log("yes")
	}
}
