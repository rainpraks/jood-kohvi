# ASCII Art reverse

Authors: Rain Praks (rainpraks) Marcus Kangur (mkangur)  Fred Kaldre (Fkaldre).

See [audit requirements]: https://github.com/01-edu/public/tree/master/subjects/ascii-art/reverse/audit

## Overview

The ASCII Art Reverse project is designed to reverse ASCII art characters based on a font file. The font file, typically named "standard.txt," contains ASCII representations of characters. The MapCharacters function reads this file and creates a mapping between each ASCII character and its corresponding representation. This mapping is stored in a map data structure (charMap). The primary goal is to enable the reversal of ASCII art, allowing the recreation of the original string from a given input file.

The program is particularly useful for scenarios where ASCII art is used as a form of encoding or representation, and there's a need to decipher or understand the encoded information. By utilizing the provided tool, users can input a file ("data.txt") containing ASCII art characters created using the mapped representations and obtain the original string.

## Usage

1. **Clone the project:** 
git clone https://01.kood.tech/git/rpraks/ascii-art-reverse

2. **Run program:** 

IMPORTANT: If you have your own ascii art text files, put them in the audit examples folder!
For reverse audit examples:
'go run . --reverse=example00.txt'<br />
'go run . --reverse=example01.txt'<br />
'go run . --reverse=example02.txt'<br />
'go run . --reverse=example03.txt'<br />
'go run . --reverse=example04.txt'<br />
'go run . --reverse=example05.txt'<br />
'go run . --reverse=example06.txt'<br />
'go run . --reverse=example07.txt'<br />

For normal (not reversed), use:
'go run . "any string"

3. **View the result:** 
The generated string will be printed to the terminal.
