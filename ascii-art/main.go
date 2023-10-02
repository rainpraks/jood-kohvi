package main

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string //

var charMap = make(map[rune]ascii)

func MapCharacters() map[rune]ascii {
	//charMap := make(map[rune]ascii)          //charMap := make(map[rune]ascii): This line creates a map where keys are Unicode characters (rune) and values are of type ascii (an array of strings representing ASCII art).
	file, err := os.ReadFile("standard.txt") //READ standart.txt file content into 'file' variable
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil

	}
	file = []byte(strings.ReplaceAll(string(file), "\r", "")) // Remove carriage return characters if any (for Windows compatibility)
	//scanner := bufio.NewScanner(bytes.NewReader(file))        // scanner that reads lines from the content. Together, this line prepares a Scanner to iterate through the lines of the file content, making it convenient for processing text files line by line.

	//teen array mis iga n\ teeb selle tükideks ja see on eraldi üksus array sees
	//mapid key on ruun ja value on need 8 rida failis
	//loop alusta spaceist ja lõpeta kuni viimase täheni. Iga loobi juures mappis key ära , läks omakorda loopi mis loopis 8 korda ja mappis values. Otsi tarka valemit. ja peale
	lines := strings.Split(string(file), "\n") // uue rea kaupa teed txt lahti.
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
	for line := 0; line < 8; line++ {
		for index, char := range inputText {
			if inputText[index] == '\\' && inputText[index+1] == 'n' {
				fmt.Print("\n") // iga value vajab \n
			}

			fmt.Print(charMap[char][line])
		}
		fmt.Print("\n")
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
