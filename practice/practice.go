package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width int
	height int
}

func (d Dimension) Area() int {
	return d.length * d.width
}

func (d Dimension) Volume() int {
	return d.length * d.width * d.height
}

func main() {
	dim := Dimension{length: 10, width: 5, height: 2}

	fmt.Println(dim.Area())
	fmt.Println(dim.Volume())
}
