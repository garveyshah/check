# ASCII-ART

***


*Project is maintained by me*

+ This project is written in Go language.

* The project uses standard Go packages.

+ Test files are meant to test the functionality of the project to ensure flawless operation of the functions

+ Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. It's Time to Write Big.

+ You can use [shadow](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt) ascii-art, [standard](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt) ascii-art, or [thinkertoy](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt) ascii-art to print out the commandline arguments.

## Usage

+ Command-line arguments are written as follows:

```bash 
go run . "example"
```


+ This project handles an input with numbers, letters, spaces, special characters and \n(new-lines).

* To know the range of ASCII characters, you can enter the following command in your terminal

```bash
man ascii
```
* man is short for manual. The command above will display a manual for ASCII characters when it runs.

+ The ASCII characters considered in this case range from ASCII character index 32 to index 126

## BANNER FORMAT *

>> Each character has a height of 8 lines

>> Characters are separated by a newline (\n)


## Script

```
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    //read the file contain the ascii-art format
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: Error Reading File", err)
		os.Exit(1)
	}

    //split the lines of the file into substrings
	input := strings.Split(string(file), "\n")

    //read the commandline arguments
	args := os.Args[1]
	for _, chrs := range args {
        //deal with cases of commandline argument not 
        // being an ASCII character
		if chrs < 32 || chrs > 126 {
			fmt.Println("Error Non Ascii Character", err)
			os.Exit(2)
		}
	}

    //deal with empty commandline argument
	if args == "" {
		return
	}

    //deal with newlines
	if args == "\n" {
		fmt.Println()
	}

    //split the commandline arguments into substrings
	argument := strings.Split(args, "\\n")
	for _, i := range argument {
		if i == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for _, chr := range i {
                    //Give the range for the specific word using
                    //the index of the first ASCII character
					startAt := int(chr-32)*9 + 1
					fmt.Print(input[startAt+j])
				}
				fmt.Println()
			}
		}
	}
}


---
---



```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
