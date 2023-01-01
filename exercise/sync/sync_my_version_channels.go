//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

const (
	DoExit = iota
	ExitOK
)

type Control int
type Job struct {
	data string
}
type Result struct {
	result int
	job    Job
}

type Count struct {
	sync.Mutex
	count int
}

func countText(jobs <-chan Job, results chan<- Result, control chan Control) {
	for {
		select {
		case job := <-jobs:
			amount := 0
			for _, r := range job.data {
				if unicode.IsLetter(r) {
					amount++
				}
			}
			results <- Result{amount, job}
		case msg := <-control:
			switch msg {
			case DoExit:
				control <- ExitOK
				return
			}
		}
	}
}

func askForText(jobs chan<- Job, control chan Control, totalCount *Count) {
	rd := bufio.NewReader(os.Stdin)
	totalCount.Lock()
	defer totalCount.Unlock()

	for {
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\nEnter text or \"q\\Q\" to exit: ")
		text, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Println()
		for _, word := range strings.Split(text[:len(text) - 1], " ") {
			jobs <- Job{word}
			if word == "q" && len(strings.Split(text[:len(text)-1], " ")) == 1 {
				control <- DoExit
				return
			}
			totalCount.count += len(word)
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan Control)
	totalCount := Count{}

	go countText(jobs, results, control)
	go askForText(jobs, control, &totalCount)

	for {
		select {
		case result := <-results:
			letterString := "letters"
			if result.result == 1 {
				letterString = "letter"
			}
			fmt.Printf("\"%v\" has %v %s in it\n", result.job.data, result.result, letterString)
		case msg := <-control:
			if msg == DoExit {
				totalCount.Lock()
				sum := totalCount.count
				totalCount.Unlock()
				fmt.Println("Exit program total character sum:", sum)
				return
			}
		}
	}
}
