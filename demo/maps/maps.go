package main

import (
	"fmt"
	"strings"
)

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	_, found := shoppingList["cereal"]
	if !found {
		var userInput string
		fmt.Print("No cereal added to shopping list, should I add it? (Y/N): ")
		fmt.Scanf("%s", &userInput)

		if strings.ToLower(userInput) == "y" {
			shoppingList["cereal"] = 1
			fmt.Println("Added cereal to your shopping list", shoppingList)
		} else {
			fmt.Println("Sure, won't add it")
		}
	}
 
	cereal := shoppingList["cereal"]

	fmt.Println("There's", cereal, "cereal in your shopping cart")

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("There're", totalItems, "in your shopping cart")
}
 