package main

import "fmt"

var number = 240

//newnumber := 280 // this is wrong since we cant do var-less initialisation outside a method

var MyVar = 40 // we have used the capital letter in the beginning of the varible hence it is a public variable
func main() {
	var username string = "Priya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of the type : %T \n", isLoggedIn)

	var smallNum uint16 = 65535
	fmt.Println("smallNum")
	fmt.Printf("The variable is of the type : %T \n", smallNum)

	var smallfloat float64 = 234.426564316047
	fmt.Println(smallfloat)
	fmt.Printf("Variable has the datatype : %T \n", smallfloat)

	// implicit variable declaration
	var emailId = "priya.gupta20044@gmail.com"
	fmt.Println(emailId)
	fmt.Printf("The type of the emailId is : %T \n",emailId);

	// var-less declaration
	// NOTE - this type of declaration only works inside a method
	// it doesn't work as a global scope

	numberOfUsers := 4924
	fmt.Println(numberOfUsers)

	fmt.Println(MyVar)
	fmt.Println(number)
}
