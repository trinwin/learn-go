package main

import "fmt"

func main() {
	switch {
		case true:
			fmt.Println(true)
		case 2 == 2:
			fmt.Println("hello")
		case false:
			fmt.Println("no print")
	}
}

