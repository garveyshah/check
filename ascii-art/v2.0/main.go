package main

import (
	"fmt"
	"os"

	"ascii/checkarguments"
	"ascii/printarguments"
)

func main() {
	// Check command-line arguments
	args, err := checkarguments.CheckArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print ASCII art for each command-line argument
	output, err := printarguments.PrintArguments(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
