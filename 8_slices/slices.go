package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "grape", "orange", "kiwi"}
	fmt.Println(fruits)
	fmt.Printf("Type of fruits is - %T\n", fruits)

	fruits = append(fruits, "mango", "peach")
	fmt.Println(fruits)

	/* Slicing */
	fruits = append(fruits[1:4])
	// Output: position 1 to position 3
	fmt.Println(fruits)

	score := make([]int, 5)
	score[0] = 88
	score[1] = 71
	score[2] = 92
	score[3] = 59
	// If less values are provided, the remaining values are set to zero
	score = append(score, 70, 82)
	fmt.Println(score)

	/* Sorting */
	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))
	// Output: true, as the slice is sorted

}
