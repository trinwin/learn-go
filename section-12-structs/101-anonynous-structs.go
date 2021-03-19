package main

import "fmt"

func main() {

	// put struct definition before the identifier
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

}
