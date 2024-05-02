package main

import (
	"strings"
	"testing"

	ascii "ascii/asciiart"
)

func TestASCIIart(t *testing.T) {
	inputStrings := []string{"H"}

	expectedOutput := `
 _    _  
| |  | | 
| |__| | 
|  __  | 
| |  | | 
|_|  |_| 
         
`

	actualOutput := ascii.ASCIIart(inputStrings)

	// Normalize line breaks for comparison
	expectedOutput = strings.ReplaceAll(expectedOutput, "\r\n", "\n")
	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n\n")

	if actualOutput != expectedOutput {
		t.Errorf("Unexpected ASCII art generated.\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}
