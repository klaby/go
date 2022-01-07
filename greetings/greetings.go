package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Wellcome!", name)
	return message
}