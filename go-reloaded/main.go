package main

import (
	"bufio"
	"fmt"
	"golang/funktsioonid"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input file> <output file>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFileContents, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	updatedContents := funktsioonid.Hex(string(inputFileContents))
	updatedContents = funktsioonid.Bin(updatedContents)
	updatedContents = funktsioonid.Up(updatedContents)
	updatedContents = funktsioonid.Low(updatedContents)
	updatedContents = funktsioonid.Cap(updatedContents)
	updatedContents = funktsioonid.UpNum(updatedContents)
	updatedContents = funktsioonid.LowNum(updatedContents)
	updatedContents = funktsioonid.CapNum(updatedContents)
	updatedContents = funktsioonid.FormatPunctuation(updatedContents)
	updatedContents = funktsioonid.FormatApostrophes(updatedContents)
	updatedContents = funktsioonid.ReplaceAWithAn(updatedContents)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Write the modified content to the output file.
	_, err = writer.WriteString(updatedContents)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	// Flush the buffered writer to ensure all data is written.
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing writer: %v\n", err)
		return
	}

}
