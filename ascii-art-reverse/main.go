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

//main.ascii [" _    _  ","| |  | | ","| |__| | ","|  __  | ","| |  | | ","|_|  |_| ","         ","         "]
//main.ascii [" _    _ ","| |  | |","| |__| |","|  __  |","| |  | |","|_|  |_|","        ","        "]

/*
matorikusu
	[        ,____    ,| ___|  ,|___ \  ,  __) | ,|____/  ,        ,        ]

	[32,45,62,84,35]
 ____
| ___|
|___ \
  __) |
|____/



lines := [][]rune{
		0     1     2    3
		{'A', 'B', ' ', ' ', 'D', ' '}, row 0
		{'A', 'E', ' ', ' ', 'D', ' '}, row 1
		{'A', 'B', ' ', ' ', 'D', ' '}, row 2
		{'A', 'B', 'C', ' ', 'D', ' '}, row 3
	}
rida := 0
	for taht:=0; taht<len(lines[0])-1;taht++{
		if lines[rida][taht]==' '{
			rida++
			taht--
		}
	}
}


map[rune]ascii

kontrollime, kas boolean foundSeperator on true: if foundSeperator{
	kui oli true, siis me loikame esimese row currentColumni indexiga katki: lines[0][:currentColumn]
	lisame katki loigatud row uude arraysse: append
	teeme seda iga rowga
	kontrollime, kas uus array vastab ascii kaardis olevate valuedega:
		kui vastab, siis lisame valuele vastavad key(t2he) uude variablei nimega outputWord
		kui ei vasta, siis teeme newArray uuesti tyhjaks
		seame boolean foundSeperatori tagasi false'iks, et loop saaks edasi minna
		returnime tagasi
loopime labi esimese row sisu, kasutame currentColumnit: for currentColumn := 0 ; i < len(row); i++ {
kui leiab tyhiku, siis teeme uue loopi, mis kontrollib kas ka teistel ridadel on samal indexil tyhik: for reaIndex := 1; reaIndex <= 8; reaIndex++ {
kui ei ole, siis laheme tagasi esimesse loopi
kui reaIndex 8'l on ka tyhik, siis me seame boolean foundSeperator to true ja naaseme tagasi algusesse
salvestab booli et column 3 oli word
siis kui oli kahendas rows samal kohal tyhik -> spliti ja appendi uute arraysse nimega newArray esimene row indexiga column-> spliti ja appendi uute arrayse teine row indexiga column
row[columnIndex] on see millega me peame tyhikuid kontrollima

kui newArray == ascii

MARCUS LAHEB SOOMA :) SEE EI OLE VEEL VALMIS AGA MIDAGI KIRJUTASIN LAHTI VAADAKE KAS ON LOOGILINE, VOID KIRJUTADA JAH




	var allSame bool
	columnIndex := 0
	currentChar := lines[rowIndex][columnIndex]
	for i := 1; i < len(matrix); i++ {
		if lines[i][columnIndex] != currentChar {
			allSame = false
			break
		}
	}








*/
