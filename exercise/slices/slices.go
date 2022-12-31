//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(title string, assemblyLine []Part) {
	fmt.Println()
	fmt.Println("---", title, "---")
	
	for _, item := range(assemblyLine) {
		fmt.Println(item)
	}
}

func main() {
	assemblyLine := []Part{"Wood", "Metal", "Planks"}
	printAssemblyLine("First assembly line", assemblyLine)

	assemblyLine = append(assemblyLine, "Structure", "Publishing")
	printAssemblyLine("Second assembly line", assemblyLine)

	assemblyLine = assemblyLine[3:]
	printAssemblyLine("Remaining assembly line", assemblyLine)
}
