package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")

	var pointer *int 
	fmt.Println("The value of ",pointer)

	myInt := 3
	refPtr := &myInt // storing the address of the variable
	// fmt.Println(refPtr,myInt)

	fmt.Println("The value of the pointer:",refPtr)
	fmt.Println("The value of the value to which the pointer points:",*refPtr)
	fmt.Printf("type is: %T \n",(refPtr))

	myStr := "Priya"
	myPtr := &myStr 

	fmt.Println(myPtr)
	fmt.Println(*myPtr)
	fmt.Printf("%T\n",myPtr)

	*refPtr = *refPtr * 2 
	fmt.Println(myInt)
}
