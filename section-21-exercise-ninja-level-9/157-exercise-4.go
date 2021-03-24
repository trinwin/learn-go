package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak()  {
	fmt.Println("I am", p.first, p.last, ". I'm", p.age)
}

func main() {
	p := person{
		"Trinh",
		"Nguyen",
		23,
	}

	p.speak()
}
