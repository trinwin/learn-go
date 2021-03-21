package main

import "fmt"

type person struct {
	first 					string
	last 						string
}

func changeMe(p *person)  {
	p.first = "Calvin"
	p.last = "Tran"
}

func main() {
	p1 := person {
		first: "Trinh",
		last: "Nguyen",
	}
	
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
