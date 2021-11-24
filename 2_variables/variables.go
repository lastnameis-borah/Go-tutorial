package main

import "fmt"

const Login_id = "admin29"

// Capital letter means it can be accessed globally

func main() {
	var user_name string = "John"
	user_name = "admin007"
	// Variable can be updated
	var user_age int = 25
	// There are various types of integer, but most of the time int datatype is used
	var user_height float32 = 5.11
	// float32 takes only upto 32-bit numbers. float64 takes upto 64-bit numbers
	// Even if the number is of higher bit, float32 convrets it to a 32 bit number
	var user_is_admin bool = true
	fmt.Println(user_name, "is", user_age, "years old and his height is", user_height, "ft.")
	fmt.Println("Is user admin? \n", user_is_admin)
	fmt.Printf("Username is of type: %T \n", user_name)
	fmt.Printf("Age is of type: %T \n", user_age)
	fmt.Printf("Height is of type: %T \n", user_height)

	fmt.Printf("\n")

	// defalut value of a int is 0
	var user_id int
	fmt.Println("User id is:", user_id)
	// defalut value of a string is empty string
	var user_spouse_name string
	fmt.Println("User name is:\n", user_spouse_name)

	// variable declaration is not mandatory, lexer takes care of it
	// But it's not a good practice
	var website = "ans.com"
	fmt.Println("Website is:", website)

	// walrus operator
	NoOfUser := 5
	fmt.Println("Number of Users:", NoOfUser)
	// walrus operator cannot be used outside a function

	fmt.Printf("\n")

	// constant
	fmt.Println("Login ID:", Login_id)
}
