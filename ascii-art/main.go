package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type ascii [8]string

func MapCharacters() map[rune]ascii {
	charMap := make(map[rune]ascii)          //charMap := make(map[rune]ascii): This line creates a map where keys are Unicode characters (rune) and values are of type ascii (an array of strings representing ASCII art).
	file, err := os.ReadFile("standard.txt") //READ standart.txt file content into 'file' variable
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil

	}
	file = []byte(strings.ReplaceAll(string(file), "\r", "")) // Remove carriage return characters if any (for Windows compatibility)
	scanner := bufio.NewScanner(bytes.NewReader(file))        // scanner that reads lines from the content. Together, this line prepares a Scanner to iterate through the lines of the file content, making it convenient for processing text files line by line.

	var currentChar rune
	var currentAscii ascii

	//prindi ascii art esimene kaheksa rida välja.
	//karakter mapi hakka 32 hakkad kõlge mappima ascii värtusi txt failist 8 kaupa
	//prindi suvaline täht, vaata kas teeb õige tähe
	//returnid mapi
	//pidu

	// see mis on allpool on puhas pask
	/*
		for scanner.Scan() { // loop iterates over each line in the file content. Initiates a loop that continues as long as there are more lines to read.
			line := scanner.Text() //Scanning each line. Inside the loop, the Text method is called to retrieve the current token, which is the text of the current line.

			if err := scanner.Err(); err != nil { //handling errors
				fmt.Println("Error reading the fail:", err)
				return nil
			}
			if line == "" { //if the line is empty, it means the ASCII representation for the current character is complete.
				charMap[currentChar] = currentAscii // store the ASCII representation in the map using the current character as the key.
				currentChar = 0                     // reset the currentchar var for the next character
				currentAscii = [8]string{}          // Reset the currentAscii array for the next character
				continue                            //	skip the rest of the loop and move to the next iteration
			}
			if currentChar == 0 { // HANDLING THE FIRST RUNE OF A CHARACTER. if currentchar is 0 it means we are starting a new character
				currentChar = []rune(line)[0] // set currentChar to the first rune of the line. Otherwise (if currentChar is not 0), it means we are continuing with the ASCII representation of the current character
			} else {
				currentAscii[len(currentAscii)-1] = line // Append each line to the corresponding element of the currentascii array.
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return nil
		}
		for k, v := range charMap {
			log.Printf("%v,%v", k, v)
		} */
	return charMap

}

func printAsciiArt(inputText string, charMap map[rune]ascii) {
	for _, char := range inputText {
		asciiArt, exists := charMap[char]
		if !exists {
			fmt.Printf("ASCII art not available for character: %c\n", char)
			continue
		}

		for _, line := range asciiArt {
			fmt.Println(line)
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
