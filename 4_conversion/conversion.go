package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your age:")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	fmt.Println("You are", age, "years old.")
	fmt.Printf("Type of age is: %T", age)

	// Conversion
	addAge, err := strconv.ParseFloat(strings.TrimSpace(age), 64)
	// Convert age string to float64 type
	// In Go, \n is added automatically at the end of the input string, so we need to trim it.

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nYou are", addAge+5, "years old.")
	}
}
