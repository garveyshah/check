package ascii

import (
	"fmt"
	"os"
	"strings"
)

// ASCIIart generates ASCII art based on the input strings
func ASCIIart(inputStrings []string) string {
	// Read the content of the ASCII art file

	asciiArtContent, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: Error Reading File", err)
		os.Exit(0)
	}
	asciiArt := strings.Split(string(asciiArtContent), "\n")

	var output strings.Builder

	// Process each input string
	for _, input := range inputStrings {
		lines := strings.Split(input, "\\n")
		for _, line := range lines {
			// Print a newline for empty lines
			if line == "" {
				output.WriteString("\n")
			} else {
				// Print ASCII art for each line
				for i := 0; i < 8; i++ {
					for _, char := range line {
						// Calculate the starting index of the character in the ASCII art
						startIndex := int(char-32)*9 + 1
						// Append the corresponding line of the ASCII art to the output
						output.WriteString(asciiArt[startIndex-1+i])
					}
					output.WriteString("\n")
				}
			}
		}
	}

	return output.String()
}

// ASCIIartContent is the ASCII art content to be used by the function
// var asciiArtContent = `
//    _____
//   / ____|
//  | |  __  __ _ _ __ ___   ___
//  | | |_ |/ _  | '_   _  \/ _ \
//  | |__| | (_| | | | | | |  __/
//   \_____|\__,_|_| |_| |_|\___|`
