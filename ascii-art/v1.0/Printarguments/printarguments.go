package asciiart

import (
	"fmt"
	"strings"
)

func PrintArguments(lineSlc []string, args []string) {
	for _, arg := range args {
		// deal with empty values and newlines
		if arg == string("") {
			continue
		}
		if arg == string("\n") {
			fmt.Println()
			return
		}

		// split the command-line argument
		lines := strings.Split(arg, "\\n")
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
}
