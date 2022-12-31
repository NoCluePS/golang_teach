package main

import "fmt"

func double(x int) int {
	return x + x;
}

func add(lhs, rhs int) int {
	return lhs + rhs;
}

func greet() {
	fmt.Println("Hello, from greet function")
}

func main() {
	greet()

	dozen := double(6);
	fmt.Printf("A dozen is %d\n", dozen)

	bakerDozen := add(dozen, 1)
	fmt.Printf("A baker's dozen is %d\n", bakerDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Printf("Another baker's dozen is %d\n", anotherBakersDozen)
}

