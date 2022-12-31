//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterate(lines []string, callBack LineCallback) {
	for _, line := range lines {
		callBack(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	sumLetters, sumDigits, sumSpaces, sumPunct := 0, 0, 0, 0

	callBack := func(line string) {
		letters, digits, spaces, punct := 0, 0, 0, 0
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++;
				sumLetters++;
			} else if unicode.IsSpace(r) {
				spaces++
				sumSpaces++;
			} else if unicode.IsPunct(r) {
				punct++
				sumPunct++
			} else if unicode.IsDigit(r) {
				digits++
				sumDigits++
			}

		}
		fmt.Printf("---\n\"%v\"\nLetters: %v\nDigits: %v\nSpaces: %v\nPunctuation: %v\n", line, letters, digits, spaces, punct)
	}
	
	lineIterate(lines, callBack)
	fmt.Printf("\n\nSum Letters: %v\nSum Digits: %v\nSum Spaces: %v\nSum Punctuation: %v\n", sumLetters, sumDigits, sumSpaces, sumPunct)
}
