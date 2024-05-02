package ascii

import (
	"fmt"
	"os"
	"strings"
)

// PrintArguments reads standard.txt and prints ASCII art for each input string
func PrintArguments(input []string) (string, error) {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	inputStrings := strings.Split(string(file), "\n")

	// strings.Builder type can be used to efficiently and easily build or concatenate strings.
	var output strings.Builder

	for _, i := range input {
		if i == "" {
			output.WriteString("\n")
		} else {
			for j := 0; j < 8; j++ {
				for _, chr := range i {
					// deal with the non-printable characters
					if chr < 32 || chr > 126 {
						return "", fmt.Errorf("Error: Non-printable character in input string")
					}
					startAt := int(chr-32)*9 + 1
					output.WriteString(inputStrings[startAt+j])
				}
				output.WriteString("\n")
			}
		}
	}
	return output.String(), nil
}
