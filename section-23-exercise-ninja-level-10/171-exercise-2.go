package main

import "fmt"

type person struct {
	first string
	last string
}

func (p* person) speak()  {
	fmt.Println("Name:", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main() {
	p1 := person {
		first: "Trinh",
		last: "Nguyen",
	}
	p2 := person {
		first: "Cuong",
		last: "Nguyen",
	}

	saySomething(&p1)
	// saySomething(p2)
}

