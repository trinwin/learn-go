package main

import (
	"fmt"

	"github.com/trinwin/go-learn/tree/main/section-27-exercise-ninja-level-12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}