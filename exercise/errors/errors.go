//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors
package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Hour = iota
	Minute
	Second
)

type Time struct {
	Hour   int
	Minute int
	Second int
}

func ParseTime(input string) (Time, error) {
	splittedTime := strings.Split(input, ":")
	var hour, minute, second int

	for i, time := range splittedTime {
		value, err := strconv.Atoi(time)
		if err != nil {
			return Time{}, fmt.Errorf("couldn't parse given time string at: %v, with index of: %v", time, i)
		}

		switch i {
		case Hour:
			hour = value
		case Minute:
			minute = value
		case Second:
			second = value
		}
	}

	return Time{hour, minute, second}, nil
}
