//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	sixSided = 6
	tenSided = 10
	twelveSided = 12
)

func rollDice(sideCount, rollCount, diceNumber int) int {
	sum := 0;

	for i := 0; i < rollCount; i++ {
		rand.Seed(time.Now().UnixNano())
		rolledSum := rand.Intn(sideCount + 1)
		fmt.Printf("\tDice #%d, on the %d roll, got %d\n", diceNumber, i+1, rolledSum)
		if rolledSum == 0 {
			sum += 1
		} else {
			sum += rolledSum
		}
	}

	return sum
}

func main() {
	var rollCount, diceCount int
	diceSideCount := tenSided
	fmt.Print("How many rolls?: ")
    _, err := fmt.Scanf("%d", &rollCount)
	if err != nil {
		fmt.Println("Input a number for roll count")
	}

	fmt.Print("How many dices?: ")
    _, err1 := fmt.Scanf("%d", &diceCount)
	if err1 != nil {
		fmt.Println("Input a number for dice count")
	}

	sum := 0
	
	for i := 0; i < diceCount; i++ {
		rolledSum := rollDice(diceSideCount, rollCount, i+1)
		fmt.Printf("\tDice #%d rolled %d\n", i+1, rolledSum)
		sum += rolledSum
	}

	fmt.Printf("Sum of all the dices is %d\n", sum)

	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	if diceCount == 2 && sum == 2 {
		fmt.Println("Snake eyes")
	}
	//  - "Lucky 7": when the total roll is 7
	if sum == 7 {
		fmt.Println("Lucky 7")
	}
	//  - "Even": when the total roll is even
	if sum % 2 == 0 {
		fmt.Println("Even")
	}else {
		fmt.Println("Odd")
	}
}
