package main

import "fmt"

func printArray(arr [5]int) {
	// array is call-by-VALUE in Go!
	// normally we don't use array in Go!
	// we use slice!
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	printArray(arr3)
}
