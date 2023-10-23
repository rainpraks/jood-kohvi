package asciiart

import (
	"fmt"
	"os"
	"strings"
)

type ascii [8]string

var charMap = make(map[rune]ascii)

func MapCharacters(fontPath string) map[rune]ascii {

	file, err := os.ReadFile(fontPath)
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
func GenerateAsciiArt(inputText string, inputFont string) string {
	charMap := MapCharacters("fonts/" + inputFont + ".txt")

	var asciiArt strings.Builder

	parts := strings.Split(inputText, "\n")
	for _, part := range parts {
		if part == "" {
			asciiArt.WriteString("\n")
			continue
		}
		for line := 0; line < 8; line++ {
			for _, char := range part {
				asciiArt.WriteString(charMap[char][line])
			}
			asciiArt.WriteString("\n")
		}
	}
	return asciiArt.String()

}
