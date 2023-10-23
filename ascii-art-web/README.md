# ASCII-art-web

Ascii-art is a program that receives a string as an input and outputs it in a graphic representation using ASCII. User can choose between three font styles: Standard, Shadow and Thinkertoy.

## Authors
Rain Praks (rpraks)

## See [audit requirements]: 
https://github.com/01-edu/public/tree/master/subjects/ascii-art-web/audit

## Usage

1. **Clone the project:** 
git clone https://01.kood.tech/git/rpraks/ascii-art-web

2. **Run program:** 
Using terminal, navigate to the main folder and execute command`go run .`

3. **View the result:** 
Open the following page in your web browser: http://localhost:8080/
There you can input your text and select a font.
To exit the program, just click CTRL + C in your terminal.

## Structure

`main.go`: Initializes and runs a web server on port 8080, serving a webpage that allows users to input text and select an ASCII font. The server then generates and displays the ASCII art representation of the input string.

`asciiart.go`: Provides the ability to build ASCII art representations of input strings and initializes ASCII art mappings from a given font file. It uses the map to translate input text into the ASCII graphics.

`webpage.html`: A simple HTML file.

`error.html`: A simple HTML file for error handling.

`standart.txt`, `shadow.txt`, and `thinkertoy.txt`: Fonts for printing ASCII art

