//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name string
	Price int
}

func getSumOfProductPrices(products [4]Product) int {
	sum := 0;
	
	for i := 0; i < len(products); i++ {
		product := products[i];
		sum += product.Price
	}

	return sum;
}

func main() {
	products := [4]Product{
		{Name: "Banana", Price: 2},
		{Name: "Mango", Price: 3},
		{Name: "Apple", Price: 4},
	}

	fmt.Println("The last item is:", products[len(products) - 2])
	fmt.Println("Total number of products:", len(products))
	fmt.Println("Total price of all products:", getSumOfProductPrices((products)))

	products[3] = Product{
		Name: "Cucumber",
		Price: 12,
	}

	fmt.Println("The last item is:", products[len(products) - 1])
	fmt.Println("Total number of products:", len(products))
	fmt.Println("Total price of all products:", getSumOfProductPrices(products))
}
