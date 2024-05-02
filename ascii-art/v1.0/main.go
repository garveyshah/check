package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	printarguments "asciiart/Printarguments"
)

func main() {
	// read the txt file
	file, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(0)
	}
	args := os.Args[1:]
	for _, arg := range args {
		for _, chr := range arg {
			if chr < 32 || chr > 126 {
				log.Fatal("Error : Non ascii/printable characters found")
				os.Exit(1)
			}
		}
	}
	if len(os.Args) < 2 {
		fmt.Println("Error : insufficient arguments")
	}
	for _, arg := range args {
		// deal with empty values and newlines
		if arg == string("") {
			return
		}
		if arg ==("\n") {
			fmt.Println()
			return
		}
	// split the txt file
	lineSlc := strings.Split(string(file), "\n")

	// output := ascii.printArguments(lineSlc)

	// print ASCII-art
	printarguments.PrintArguments(lineSlc, args)
}
}
