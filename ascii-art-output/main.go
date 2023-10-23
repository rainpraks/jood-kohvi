package main

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func main() {
	var outputFileName string
	var inputText string
	var inputFont string

	// go run . --output=<fileName.txt> "something" standard
	// go run . "hello" standard
	// go run . "hello"
	if len(os.Args) == 4 {
		outputFileName = os.Args[1]
		inputText = os.Args[2]
		inputFont = os.Args[3]
	} else if len(os.Args) == 3 && (os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy") {
		inputText = os.Args[1]
		inputFont = os.Args[2]
	} else if len(os.Args) == 2 {
		inputText = os.Args[1]
		inputFont = "standard"
	} else {
		fmt.Printf("Error, missing or too many arguments. Examples:\n--output=<fileName.txt> \"something\" standard\n\"hello\" standard\n\"hello\"\n")
		return
	}

	//ascii style file we make the map of
	styleFile, err := os.ReadFile("fonts/" + inputFont + ".txt")

	if err != nil {
		fmt.Println("Error opening styles file. Make sure the font is written in lowercase.", err)
		return
	}
	asciiMap := MapCharacters(styleFile)

	if len(os.Args) > 4 {
		fmt.Println("Error reading input, too many args")
		return
	}

	var newName string
	//checks for --output= flag
	if len(outputFileName) > 9 && outputFileName[:9] == "--output=" {
		newName = outputFileName[9:]
		outputFile, err := os.Create(newName)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		printAsciiArtToOutput(inputText, asciiMap, outputFile)
		return
	}
	printAsciiArt(inputText, asciiMap)

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

// 1.Takes input text and the character map as parameters.
// 2.Splits the input text into parts based on the "\n" delimiter.
// 3.Prints the ASCII art representation of each character in the input text.
func printAsciiArtToOutput(inputText string, charMap map[rune]ascii, outputFile *os.File) {
	parts := strings.Split(inputText, "\\n")
	for _, part := range parts {
		if part == "" {
			outputFile.WriteString("")
			continue
		}
		for line := 0; line < 8; line++ {
			for _, char := range part {
				outputFile.WriteString(charMap[char][line])
			}
			outputFile.WriteString("\n")
		}
	}
}
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
