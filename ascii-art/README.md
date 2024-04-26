** ASCII-ART**

*Project is maintained by me*

+ This project is written in Go language

+ The project uses standard Go packages

+ Test files are meant to test the functionality of the said project to ensure flawless operation of the functions

+ Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

+ Below is a representation of an ASCII character

@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$


+ This project handles an input with numbers, letters, spaces, special characters and \n(new-lines).

+ The ASCII characters considered in this case range from ASCII character index 32 to index 126

* BANNER FORMAT *

>> Each character has a height of 8 lines
>> Characters are separated by a newline (\n)



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

```

