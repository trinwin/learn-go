package main

import "fmt"

func main() {
	bar("Trinh")
	s1 := woo("Kem")
	fmt.Println(s1)
}

// func (r receiver) identifier (parameters) (return) {code}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from ", st)
}
