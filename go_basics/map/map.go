package main

import "fmt"

func main() {
	m := map[string]string{"name": "kky", "major": "computer science"}
	m2 := make(map[string]int) // = empty map
	var m3 map[string]int      // = nil, which is an empty map

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
