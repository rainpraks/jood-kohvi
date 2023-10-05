package main

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func main() {
	styleFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error opening styles file:", err)
		return
	}
	styleFile = []byte(strings.ReplaceAll(string(styleFile), "\r", ""))
	asciiMap := MapCharacters(styleFile)

	inputFile, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	yuh(inputFile, asciiMap)
}

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

func yuh(inputFile []byte, asciiMap map[rune]ascii) {
	var jahOn bool
	var newValue ascii
	lastIndex := 0
	lines := strings.Split(string(inputFile), "\n")

	strings := 0
	for letters := 0; letters < len(lines[0]); letters++ {
		if lines[strings][letters] == ' ' {
			jahOn = true
			strings++
			letters--
		} else {
			jahOn = false
			strings = 0
		}

		if strings == 7 && jahOn {
			letters++
			for i := 0; i < 8; i++ {
				newValue[i] = lines[i][lastIndex : letters+1]
			}
			for k, value := range asciiMap {
				if value == newValue {
					fmt.Print(string(k))
					lastIndex = letters + 1
				}
			}
			strings = 0
		}
	}
}
