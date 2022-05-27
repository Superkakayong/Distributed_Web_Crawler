package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	// slice in Go is call-by-REFERENCE!
	// because slice is just a view of the array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(arr)

	// reslice
	// can only reslice backward, not forward
	// s[i] cannot go beyond len(s); reslice cannot go beyond cap(s)
}
