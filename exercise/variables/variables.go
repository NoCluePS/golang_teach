//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var colour = "Blue"
	birthYear, ageYears := 2006, 16
	var (
		firstInitial = "P"
		secondInitial = "S"
	)
	var ageInDays int
	ageInDays = ageYears*365

	fmt.Printf("Colour = %s\nBirthYear = %d\nAge in years = %d\nFirst initial = %s\nLast initial = %s\nAge in days = %d\n", colour, birthYear, ageYears, firstInitial, secondInitial, ageInDays)
}
