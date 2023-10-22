package main

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func main() {
	//ascii style file we make the map of
	styleFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error opening styles file:", err)
		return
	}
	styleFile = []byte(strings.ReplaceAll(string(styleFile), "\r", "")) //replace \r with nothing. For windows system.
	asciiMap := MapCharacters(styleFile)

	var newName string
	inputFileName := os.Args[1]

	if len(os.Args) > 2 {
		fmt.Println("Error reading input, too many args")
		return
	}

	//checks for --reverse= flag
	if len(inputFileName) > 10 && inputFileName[:10] == "--reverse=" {
		newName = inputFileName[10:]
		inputFile, err := os.ReadFile("audit examples/" + newName)
		if err != nil {
			fmt.Printf("Error reading --reverse file, no files named %v: ", newName)
		}
		Reverse(inputFile, asciiMap)
		return
	}
	printAsciiArt(inputFileName, asciiMap)

}

// 1. Reads a font file (e.g., "standard.txt") containing ASCII art characters.
// 2. Maps each character to its corresponding ASCII representation.
// 3. Stores the mappings in a map data structure.
func MapCharacters(file []byte) map[rune]ascii {
	lines := strings.Split(string(file), "\n")
	for char := ' '; char <= '~'; char++ {
		asciimap := ascii{}
		for line := 0; line < 8; line++ {
			asciimap[line] = lines[1+line+int(char-' ')*9]
		}
		charMap[char] = asciimap
	}
	return charMap
}

// 1. Process input line by line.
// 2. Identify sequences corresponding to mapped ASCII representations.
// 3. Look up and print the original character.
// 4. Repeat until the entire input is processed.
func Reverse(inputFile []byte, asciiMap map[rune]ascii) {
	lines := strings.Split(string(inputFile), "\n")

	var foundWhitespace bool
	var newValue ascii
	lastIndex := 0
	row := 0

	for column := 0; column < len(lines[0]); column++ {
		if lines[row][column] == ' ' { //keeps the loop on same column, but changes the row, if it finds whitespace
			foundWhitespace = true
			row++
			column--
		} else {
			foundWhitespace = false
			row = 0
		}

		if row == 8 && foundWhitespace {
			column++
			for i := 0; i < 8; i++ { //slices each line at whitespace column and adds to newValue
				newValue[i] = lines[i][lastIndex : column+1]
			}
			for key, value := range asciiMap { //compares asciiMap value with our newValue
				if value == newValue {
					fmt.Print(string(key))
					lastIndex = column + 1
				}
			}
			row = 0
		}
	}
	fmt.Println()
}

// 1.Takes input text and the character map as parameters.
// 2.Splits the input text into parts based on the "\n" delimiter.
// 3.Prints the ASCII art representation of each character in the input text.
func printAsciiArt(inputText string, charMap map[rune]ascii) {
	parts := strings.Split(inputText, "\\n")
	for _, part := range parts {
		if part == "" {
			fmt.Println("")
			continue
		}
		for line := 0; line < 8; line++ {
			for _, char := range part {

				fmt.Print(charMap[char][line])
			}
			fmt.Print("\n")
		}
	}
}
