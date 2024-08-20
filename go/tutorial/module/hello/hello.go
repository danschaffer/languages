package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Joe")
	fmt.Println(message)
	fmt.Println(greetings.Hello("Suz"))
}
