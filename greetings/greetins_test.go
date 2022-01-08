package greetings

import (
	"regexp"
	"testing"
)

// Must call the Hello function passing a valid argument and return a valid hello message.
func TestHelloName(t *testing.T) {
	name := "hiukky"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("hiukky")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("hiukky") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Should return an error for an invalid argument
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}