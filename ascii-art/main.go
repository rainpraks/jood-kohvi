package main

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func MapCharacters() map[rune]ascii {

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil

	}
	file = []byte(strings.ReplaceAll(string(file), "\r", ""))

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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"inputtext\"")
		return
	}
	inputText := os.Args[1]

	charMap := MapCharacters()

	printAsciiArt(inputText, charMap)
}
