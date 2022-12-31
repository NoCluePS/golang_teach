package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (c *Stuff) Get(index int) (int, error) {
	if index > len(c.values) {
		return 0, fmt.Errorf("no element at index %v", index)
	}

	return c.values[index], nil
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value is", value)
	}
}
