package main

import "fmt"

// modules are imported automatically one the code is executed
// fmt is a package that is used to print text to the screen

func main() {
	fmt.Println("Hello, World!")
	// Uppercase depicts that the method from fmt package is exported publically
	// Almost everything is type (string, bool, integer, floating, complex, etc..) in Go
}

// Lexer adds a semicolon in the end and also make grammatical corrections
// But it doesn't make logical corrections
