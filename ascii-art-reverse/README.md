# ASCII Art reverse

Authors: Rain Praks (rainpraks) Marcus Kangur (mkangur)  Fred Kaldre (Fkaldre).

See [task description]: exercise.txt [audit requirements]: https://github.com/01-edu/public/tree/master/subjects/ascii-art/reverse/audit

## Overview

The ASCII Art Reverse project is designed to reverse ASCII art characters based on a font file. The font file, typically named "standard.txt," contains ASCII representations of characters. The MapCharacters function reads this file and creates a mapping between each ASCII character and its corresponding representation. This mapping is stored in a map data structure (charMap). The primary goal is to enable the reversal of ASCII art, allowing the recreation of the original string from a given input file.

The program is particularly useful for scenarios where ASCII art is used as a form of encoding or representation, and there's a need to decipher or understand the encoded information. By utilizing the provided tool, users can input a file ("data.txt") containing ASCII art characters created using the mapped representations and obtain the original string.

## Usage

1. **Clone the project:** 
clone git https://01.kood.tech/git/rpraks/ascii-art-reverse

2. **Run program:** 
go run . "data.txt" .Replace data.txt content with audits question

3. **View the result:** 
The generated string will be printed to the terminal.


## Description of the Code

MapCharacters Function:
1. Reads a font file (e.g., "standard.txt") containing ASCII art characters.
2. Maps each character to its corresponding ASCII representation.
3. Stores the mappings in a map data structure.

yuh function:
1. Process input line by line.
2. Identify sequences corresponding to mapped ASCII representations.
3. Look up and print the original character.
4. Repeat until the entire input is processed.

Main Function:
1. Read font file using MapCharacters to establish mappings.
2. Read input file and invoke yuh for the reversal process.
3. Print the generated result (original string) to the terminal.
