package main

import (
	"reflect"
	"testing"
)

func TestMapCharacters(t *testing.T) {
	tests := []struct {
		name string
		want map[rune]ascii
	}{
		// TODO: Add test cases.
		{"test1", map[rune]ascii{'h': ascii{"", " _", "| |", "| |_", "| '_ \\", "| | | |", "|_| |_|"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapCharacters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
