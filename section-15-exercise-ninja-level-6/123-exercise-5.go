package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * 2
}

// make sure to add full method definition
type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type){
		case square:
			fmt.Println("area of a square:", s.(square).area())
		case circle:
			fmt.Println("area of a circle:", s.(circle).area())
	}
}

func main() {
	s := square {
				length: 4,
	}

	c := circle {
				radius: 4,
	}

	info(s)
	info(c)
}
