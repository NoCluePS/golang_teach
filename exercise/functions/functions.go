//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s\n", name);
}

func getMessage() string {
	return "Greetings from sendMessage()";
}

func add(a, b, c int) int {
	return a + b + c;
} 

func getNumber() int {
	return 123;
}

func getTwoNumbers() (int, int) {
	return 123, 321
}

func main() {
	var name string;
	fmt.Print("What's your name?: ")
	fmt.Scanln(&name)
	greet(name);

	fmt.Println(getMessage())

	a, b := getTwoNumbers()
	addedNumber := add(getNumber(), a, b)
	fmt.Printf("Added number is: %d\n", addedNumber)

	fmt.Printf("Number got: %d\n", getNumber())

	fmt.Printf("First got number: %d\n",a )
}
