package main

import "fmt"

func main() {
	var fruits [3]string
	// No. of elements and type of elements are to defined previously
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	// If lesser no. of elements are given, then a space is added to the end of the array
	fmt.Println(fruits)
	// Output: [Apple Orange  ]
	fmt.Println(len(fruits))
	// Output: 3, as 3 locations are already allocated for the array

	// Better way of defining array
	var fruits1 = [2]string{"Mango", "Strawberry"}
	fmt.Println(fruits1)
}
