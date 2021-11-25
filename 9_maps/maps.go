package main

import "fmt"

func main() {
	m := make(map[string]int)
	// (map[Type of key]Type of value)
	m["k1"] = 7
	m["k2"] = 9
	m["k3"] = 11
	fmt.Println("The map is -", m)
	fmt.Println("First key value is -", m["k1"])
	fmt.Println()

	// Delete
	delete(m, "k2")
	fmt.Println("The map is -", m)

	// for loops
	for key, value := range m {
		fmt.Printf("The key is %v and the value is %v\n", key, value)
		// %v is the format specifier for value
	}
}
