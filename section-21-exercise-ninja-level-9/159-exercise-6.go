package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
}
