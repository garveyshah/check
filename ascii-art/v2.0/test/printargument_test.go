package printargument_test

import (
	"testing"
)

func TestPrintArguments(t *testing.T) {
	// Test case with valid input
	input := []string{"Hello", "World"}
	expected := "  _   _      _ _        __        __         _     _ \n | | | | ___| | | ___   \\ \\      / /__  _ __| | __| |\n | |_| |/ _ \\ | |/ _ \\   \\ \\ /\\ / / _ \\| '__| |/ _` |\n |  _  |  __/ | | (_) |   \\ V  V / (_) | |  | | (_| |\n |_| |_|\\___|_|_|\\___( )   \\_/\\_/ \\___/|_|  |_|\\__,_|\n                     |/                              \n"
	output, err := printarguments.PrintArguments(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if output != expected {
		t.Errorf("Output mismatch. Expected:\n%s\nGot:\n%s\n", expected, output)
	}

	// Test case with invalid character in input
	input = []string{"Hello", "123"}
	_, err = printarguments.PrintArguments(input)
	if err == nil {
		t.Errorf("Expected error for invalid character in input")
	}

	// Test case with empty input
	input = []string{}
	expected = ""
	output, err = printarguments.PrintArguments(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if output != expected {
		t.Errorf("Output mismatch. Expected:\n%s\nGot:\n%s\n", expected, output)
	}
}
