package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, element := range(slice) {
		fmt.Println(element, "is at index", i)
		for _, ch := range(element) {
			fmt.Printf("   %q\n", ch)
		}
	}
}
