package main

import (
	"fmt"
	"math"
)

// CANNOT use colon to define variables outside a function
// these are not "global" variables. They are just variables within a package
var (
	aa = 3
	ss = "kkk"
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d, %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "KKY"
	fmt.Println(a, b, s)
}

func variableWithNoTypes() {
	var a, b, s = 3, 4, "KKY" // can write variables with different types in the same line
	//var s = "KKY"
	fmt.Println(a, b, s)
}

func variableSimplerWay() {
	a, b, s := 3, 4, "KKY" // use colon instead of "var". BETTER WAY! DON'T USE var!
	b = 5
	fmt.Println(a, b, s)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func main() {
	//fmt.Println("hello")
	//fmt.Println(runtime.GOARCH) // check your code is running on what architecture
	variableZeroValue()
	variableInitialValue()
	variableWithNoTypes()
	variableSimplerWay()
	fmt.Println(aa, ss)
	consts()
}
