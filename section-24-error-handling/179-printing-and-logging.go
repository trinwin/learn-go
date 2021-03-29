package main

import (
	"os"
)

// Examples here: https://github.com/GoesToEleven/go-programming/tree/master/code_samples/006-error-handling

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}
}

// http://godoc.org/builtin#panic
// Println calls Output to print to the standard logger. 
// Arguments are handled in the manner of fmt.Println.


/*
Package log implements a simple logging package ... 
writes to standard error and prints the date and time of each logged message
*/

