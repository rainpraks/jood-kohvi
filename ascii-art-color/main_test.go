package main

import "testing"

func Test_printAsciiArt(t *testing.T) {
	type args struct {
		inputText string
		color     string
		letters   string
		charMap   map[rune]ascii
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// add args and parameters here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printAsciiArt(tt.args.inputText, tt.args.color, tt.args.letters, tt.args.charMap)
		})
	}
}
