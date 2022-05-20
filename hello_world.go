package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello")
	fmt.Println(runtime.GOARCH) // check your code is running on what architecture
}
