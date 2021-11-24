package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	/* New reader is created using the bufio pkg and
	the reader is reading from standard input output (Stdin) using os pkg*/
	fmt.Println("Enter your name:")

	// comma error method
	name, _ := reader.ReadString('\n') // read string until new line
	// name: Input
	// _: error
	/* (In many cases, the error is ignored and marked with "_".
	But if there's any, then it is mentioned as "err")*/

	fmt.Println("Hello", name)
	fmt.Printf("Type of name is: %T", name)
	// output: Everything is stored as string when using user input
}
