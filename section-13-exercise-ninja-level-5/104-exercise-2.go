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
		lastName: "Tran",
		favIceCreamFlavors: []string {"vanila", "mint"},
	}

	m := map[string]person {
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, flavor := range v.favIceCreamFlavors {
			fmt.Println(flavor)
		}
	}
}
