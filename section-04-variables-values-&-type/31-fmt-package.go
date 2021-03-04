package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // type of value
	fmt.Printf("%b\n", y)  // base 2 binary
	fmt.Printf("%x\n", y)  // base 16 hex
	fmt.Printf("%#x\n", y) // base 17 hex - alternate format: add leading 0b for binary (%#b), 0 for octal (%#o),
	//0x or 0X for hex (%#x or %#X); suppress 0x for %p (%#p);
	y = 911
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
}
