package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func validate(message interface {}, err error) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func printFail() {
	message, err := greetings.Hello("")
	validate(message, err)
}

func printOne() {
	message, err := greetings.Hello("Hiukky")
	validate(message, err)
}

func printMany() {
	messages, err := greetings.Hellos([]string{"Hiukky", "Romullo", "Devmarsh"})
	validate(messages, err)
}

func main() {
	printOne()
	printMany()
	printFail()
}