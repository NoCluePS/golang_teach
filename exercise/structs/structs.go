//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func getPerimeter(rect Rectangle) int {
	perimeter := 2 * (rect.height + rect.width)
	return perimeter
}

func getArea(rect Rectangle) int {
	area := rect.height * rect.width
	return area
}

func main() {
	var width, height int
	fmt.Print("Give the width of rect: ")
	if _, err := fmt.Scanf("%d", &width); err != nil {
		fmt.Println("Input a number for width")
	}
	fmt.Print("Give the height of rect: ")
	if _, err := fmt.Scanf("%d", &height); err != nil {
		fmt.Println("Input a number for height")
	}

	rectangle := Rectangle{
		width,
		height,
	}

	fmt.Println("The perimeter of the rect is", getPerimeter(rectangle))
	fmt.Println("The area of the rect is", getArea(rectangle))

	rectangle.height = rectangle.height*2
	rectangle.width = rectangle.width*2

	fmt.Println("The perimeter of the rect*2 is", getPerimeter(rectangle))
	fmt.Println("The area of the rect*2 is", getArea(rectangle))
}
