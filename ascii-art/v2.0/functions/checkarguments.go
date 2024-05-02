package ascii

import (
	"fmt"
	"os"
)

// CheckArguments checks if the command-line arguments are valid
func CheckArguments() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("error: insufficient number of arguments")
	}

	var args []string

	if args == "" {
		return
	}
	if args == "\n" {
		fmt.Println()
		os.Exit(0)
	}

	for _, arg := range os.Args[1:] {
		for _, ch := range arg {
			if ch < 32 || ch > 126 {
				return nil, fmt.Errorf("Error: Non-printable character in input string")
			}
		}
		args = append(args, arg)
	}

	return args, nil
}
