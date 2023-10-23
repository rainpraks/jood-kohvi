package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func main() {
	// reads font file
	styleFile, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	//compatibility for windows os
	charMap := MapCharacters(styleFile)

	var inputText string
	var colorInput string
	var letters string

	colorArgs := os.Args[1]

	if len(os.Args) > 4 {
		fmt.Println("Error reading input, too many args")
		return
	}
	// detects designated colors and declare it as new variable
	if len(colorArgs) > 8 && colorArgs[:8] == "--color=" {
		colorInput = colorArgs[8:]
	}
	// call for coloValitator (see line 111)
	color, err := colorValidator(colorInput)
	if err != nil {
		fmt.Printf("%v is not a color supported in this program", color)
		return

	}
	//if no desired letters to be colored then color full "inputText"
	if len(os.Args) == 3 {
		letters = ""
		inputText = os.Args[2]
		// reads designated letters and inputText
	} else if len(os.Args) == 4 {
		letters = os.Args[2]
		inputText = os.Args[3]
	}

	printAsciiArt(inputText, color, letters, charMap)
}

// creates the map for ascii symbols
func MapCharacters(styleFile []byte) map[rune]ascii {

	lines := strings.Split(string(styleFile), "\n")
	for char := ' '; char <= '~'; char++ {
		asciimap := ascii{}
		for line := 0; line < 8; line++ {
			asciimap[line] = lines[1+line+int(char-' ')*9]
		}
		charMap[char] = asciimap
	}

	return charMap

}

// prints out designated letters or full inputText string
func printAsciiArt(inputText string, color string, letters string, charMap map[rune]ascii) {
	parts := strings.Split(inputText, "\\n")
	for _, part := range parts {
		if part == "" {
			fmt.Println("")
			continue
		}
		for line := 0; line < 8; line++ {
			for _, char := range part {

				if letters == "" {
					fmt.Print(color, charMap[char][line], "\033[0m") // Reset color
				} else {
					isLetterInInputtext := false
					for _, ch := range letters {
						if ch == char {
							isLetterInInputtext = true
						}
					}

					if isLetterInInputtext {
						fmt.Print(color, charMap[char][line], "\x1b[0m") // Reset color
					} else {
						fmt.Print(charMap[char][line])
					}

				}

			}
			fmt.Print("\n")
		}
	}
}

// color validator checks available colors and adds them to the variable
func colorValidator(colorInput string) (string, error) {
	const (
		RED     = "\033[31m"
		GREEN   = "\033[32m"
		BLUE    = "\033[34m"
		YELLOW  = "\033[33m"
		WHITE   = "\033[37m"
		ORANGE  = "\033[38;2;255;165;0m"
		BLACK   = "\u001b[30m"
		MAGENTA = "\u001b[35m"
		CYAN    = "\u001b[36m"
	)

	switch colorInput {
	case "red":
		return RED, nil
	case "green":
		return GREEN, nil
	case "blue":
		return BLUE, nil
	case "yellow":
		return YELLOW, nil
	case "white":
		return WHITE, nil
	case "orange":
		return ORANGE, nil
	case "black":
		return BLACK, nil
	case "magenta":
		return MAGENTA, nil
	case "cyan":
		return CYAN, nil
	default:
		return colorInput, errors.New("error")
	}
}
