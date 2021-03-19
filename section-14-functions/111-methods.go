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
	fmt.Println("I am", s.first, s.last)
	
}

func main() {

	sa1 := secretAgent{
		person: person{
							"Trinh",
							"Nguyen",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()


}
