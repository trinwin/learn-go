package main

import "fmt"

type person struct {
	firstName 					string
	lastName 						string
	favIceCreamFlavors 	[]string
}

func main() {
	p1 := person {
		firstName: "Trinh",
		lastName: "Nguyen",
		favIceCreamFlavors: []string {"cookie and cream", "chocolate"},
	}
	p2 := person {
		firstName: "Cuong",
		lastName: "Nguyen",
		favIceCreamFlavors: []string {"vanila", "mint"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favIceCreamFlavors {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.favIceCreamFlavors {
		fmt.Println(v)
	}
}
