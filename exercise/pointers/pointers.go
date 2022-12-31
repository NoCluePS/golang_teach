//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag SecurityTag
}

func LogItems(items []Item) {
	fmt.Println()
	for _, item := range items {
		fmt.Println(item)
	}
}

func Activate(tag *SecurityTag) {
	*tag = Active
}

func Deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func Checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		Deactivate(&items[i].tag)
	}
}

func main() {
	shoppingItems := []Item{{"Mangos", Active}, {"Bananas", Active}, {"Beer", Active}, {"Apples", Active}}
	LogItems(shoppingItems)
	Deactivate(&shoppingItems[2].tag)
	LogItems(shoppingItems)
	Checkout(shoppingItems)
	LogItems(shoppingItems)
}
