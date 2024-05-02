package main

import (
	"fmt"
	"os"

	ascii "ascii/asciiart"
	//"github.com/Murzuqisah/ascii/ascii" // Import the ascii package from your module
)

func main() {
	// Define input strings
	inputStrings := os.Args[1:]

	// Generate ASCII art
	asciiArt := ascii.ASCIIart(inputStrings)

	// Print the generated ASCII art
	fmt.Println(asciiArt)
}
