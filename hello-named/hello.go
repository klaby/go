package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Hiukky")
	fmt.Println(message)
}