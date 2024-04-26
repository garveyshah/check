package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: Error Reading File", err)
		os.Exit(1)
	}

	input := strings.Split(string(file), "\n")

	args := os.Args[1]
	for _, chrs := range args {
		if chrs < 32 || chrs > 126 {
			fmt.Println("Error Non Ascii Character", err)
			os.Exit(2)
		}
	}

	if args == "" {
		return
	}

	if args == "\n" {
		fmt.Println()
	}

	argu := strings.Split(args, "\\n")
	for _, i := range argu {
		if i == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for _, chr := range i {
					startAt := int(chr-32)*9 + 1
					fmt.Print(input[startAt+j])
				}
				fmt.Println()
			}
		}
	}
}
