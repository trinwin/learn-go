package main

import "fmt"

func main() {
	kem := struct {
		breed string
		age int
		color string
	}{
		breed: "siamese",
		age: 2,
		color: "cream",
	}

	fmt.Println(kem)

}
