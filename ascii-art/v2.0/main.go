package main

import (
	"fmt"
	"os"

	ascii "ascii/functions"
)

func main() {
	// Check command-line arguments
	args, err := ascii.CheckArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print ASCII art for each command-line argument
	output, err := ascii.PrintArguments(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
