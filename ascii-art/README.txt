# ASCII Art

Authors: Rain Praks (rainpraks) and Fred Kaldre (Fkaldre).

See [task description]: exercise.txt [audit requirements]: https://github.com/01-edu/public/tree/master/subjects/ascii-art/audit.

## Overview

This Go program converts input text into ASCII art using predefined character maps. The character maps are loaded from an external file, and the resulting ASCII art is printed to the terminal.

## Usage

1. **Clone the project:** 
clone git https://01.kood.tech/git/rpraks/ascii-art

2. **Run program:** 
go run . "writetexthere" .Replace "writetexthere" with your desired text.

3. **View the result:** 
The generated ASCII art will be printed to the terminal.

4. **Change the font:**
If you want to use a different font, modify the font file in the `main.go` file. Change the filename on line 15:
file, err := os.ReadFile("standard.txt") // Change "standard.txt" to the desired font file

Description of the Code

MapCharacters Function:
1.Reads a font file (e.g., "standard.txt") containing ASCII art characters.
2.Maps each character to its corresponding ASCII representation.
3.Stores the mappings in a map data structure.

printAsciiArt Function:
1.Takes input text and the character map as parameters.
2.Splits the input text into parts based on the "\n" delimiter.
3.Prints the ASCII art representation of each character in the input text.

Main Function:
1.Checks the command-line arguments. Exits if the number of arguments is not equal to 2.
2.Reads the input text from the command line.
3.Calls MapCharacters to load the character mappings.
4.Calls printAsciiArt to print the ASCII art representation of the input text.
