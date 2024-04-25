package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read the txt file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	// split the txt file
	lineSlc := strings.Split(string(file), "\n")

	// deal with empty values and newlines
	if os.Args[1] == "" {
		return
	}
	if os.Args[1] == "\n" {
		fmt.Println()
		return
	}
	if len(os.Args) > 2 {
		return
	}
	// fmt.Printf("%q\n", os.Args[1])

	// split the command-line argument
	lines := strings.Split(string(os.Args[1]), "\\n")
	for _, value := range lines {
		// print newline for empty value
		if value == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, letter := range value {
					startIndex := int(letter-32)*9 + 1
					fmt.Print(lineSlc[startIndex+i])

				}
				fmt.Println()
			}
		}
	}
}
