package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8;

	if quiz1 > quiz2 {
		fmt.Printf("quiz1 scored higher than quiz2\n")
	} else if quiz1 < quiz2 {
		fmt.Printf("quiz2 scored higher than quiz1\n")
	} else {
		fmt.Printf("quiz1 & quiz2 have the same score\n")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Printf("acceptable grades\n")
	} else {
		fmt.Println("unacceptable grades")
	}
}
