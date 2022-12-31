package main

import "fmt"

func main() {
	var myName = "Pijus"
	fmt.Printf("my name is %s\n", myName)

	var name string = "Kathy"
	fmt.Printf("name = %s\n", name)

	userName := "admin"
	fmt.Printf("username = %s\n", userName)

	var sum int
	fmt.Printf("The sum is %d\n", sum)

	part1, other := 1, 5
	fmt.Printf("part 1 is %d, other is %d\n", part1, other)

	part2, other := 2, 0
	fmt.Printf("part 2 is %d, other is %d\n", part2, other)

	sum = part1 + part2
	fmt.Printf("sum = %d\n", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Printf("lessonName = %s\nlessonType = %s\n", lessonName, lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Printf("%s %s\n", word1, word2)
}
