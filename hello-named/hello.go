package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func validate(message string, err error) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func main() {
	// Success
	message, err := greetings.Hello("Hiukky")
	validate(message, err)

	// Fail
	message, err = greetings.Hello("")
	validate(message, err)
}