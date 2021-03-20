package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person 
	ltk bool
}

// attach this function to any values of type secretAgent
// function of an object
// attaching a function to a type with a RECEIVER
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secret agent speaks")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speaks")
}

// Interface - allows a value to be of more than one type
type human interface {
	speak() // if you have this method then you are my type
}

type hotdog int

func bar(h human) {
	switch h.(type){
		case person:
			fmt.Println("I was passed into bar (person)", h.(person).first) // (person) is assertion
		case secretAgent:
			fmt.Println("I was passed into bar (secretAgent)", h.(secretAgent).first)
	}
}

func main() {
	// sa1 is of type secretAgent and human
	sa1 := secretAgent{
		person: person{
							"Trinh",
							"Nguyen",
		},
		ltk: true,
	}

	p1 := person {
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(p1)

	// Polymorphism - This function can take in different types
	bar(sa1)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// ASSERTION

}
