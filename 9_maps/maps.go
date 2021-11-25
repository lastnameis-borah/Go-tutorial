package main

import "fmt"

func main() {
	m := make(map[string]int)
	// (map[Type of key]Type of value)
	m["k1"] = 7
	m["k2"] = 9
	fmt.Println(m)
}
