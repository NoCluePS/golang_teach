//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Stats struct {
	nonBlank, commands int
}

func main() {
	r := bufio.NewReader(os.Stdin)

	stats := Stats{
		nonBlank: 0,
		commands: 0,
	}
	fmt.Println("Type q\\Q to exit the application")
	for {
		input, inputErr := r.ReadString('\n')
		n := strings.TrimSpace(input)
		n = strings.ToLower(n)
		if n == "" {
			continue
		} else {
			stats.nonBlank++
		}

		if n == "hello" {
			stats.commands++;
			fmt.Println("Hello there, dear stranger!")
		} else if n == "bye" {
			stats.commands++;
			fmt.Println("Well, bye then, dear stranger!")
		} else if n == "q" || inputErr == io.EOF {
			break
		} else {
			fmt.Println("Not an implemented command yet, sorry :D")
		}

		if inputErr != nil {
			fmt.Println(inputErr)
		}
	}

	fmt.Printf("Command count: %v\nNon-blank lines: %v\n", stats.commands, stats.nonBlank)
}
