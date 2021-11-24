package main

import "fmt"

func main() {
	var ptr *int
	// *int - It is a pointer type with base int
	// * indicates that ptr is a pointer to an int
	fmt.Println("Value of default pointer is -", ptr)
	fmt.Println()

	num := 10
	var ptr1 = &num
	// ptr1 is a pointer referencing to the address of num
	// &num returns the address of num
	fmt.Println("Value of ptr1 pointer is -", ptr1)
	// output: Memory address of num.
	fmt.Println("Value of ptr1 pointer is -", *ptr1)
	// *ptr1 is called dereferencing as it gives the value from the address
	fmt.Println()

	*ptr1 = *ptr1 + 10
	fmt.Println("Value of new pointer is -", *ptr1)
	fmt.Println("Value of new pointer is -", num)
	// output: 20; This proves that the change occur at the variable itself and not in the copy of the variable.
}

// If a variable is given, the a "&" is added to the variable name to get the address of the variable.
// If a memory address of a variable is given, the "*" is added to the memory address to get the value of the variable.

// Manipulating the variable using the memory location is much more efficient than using the copy of the variable to change it's value every time we use it
