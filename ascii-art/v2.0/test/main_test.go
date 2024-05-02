package main

import (
	"os"
	"testing"

	"ascii/checkarguments"
)

func TestCheckArguments(t *testing.T) {
	// Test case with valid input
	os.Args = []string{"program", "Hello"}
	_, err := checkarguments.CheckArguments()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case with invalid number of arguments
	os.Args = []string{"program"}
	_, err = checkarguments.CheckArguments()
	if err == nil {
		t.Errorf("Expected error for incorrect number of arguments")
	}

	// Test case with invalid character in input
	os.Args = []string{"program", "123"}
	_, err = checkarguments.CheckArguments()
	if err == nil {
		t.Errorf("Expected error for invalid character in input")
	}

	// Test case with invalid characters in input
	input := []string{"Line\twith\ttab", "Newline\\n", "Space\\s", "CarriageReturn\\r", "VerticalTab\\v", "FormFeed\\f", "Backspace\\b"}
	_, err = checkarguments.CheckArguments(input)
	if err == nil {
		t.Errorf("Expected error for invalid characters in input")
	}
}
