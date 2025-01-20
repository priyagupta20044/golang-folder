package main

import "fmt"

func main() {
	fmt.Println("give")
	defer fmt.Println("Away")
	fmt.Println("hello")
	defer fmt.Println("okay")
	defer fmt.Println("wordle")
	fmt.Println("charger")

	mydef()
}
// Stack ->  0 1 2 3 4 5 : Here, there is a defer statement in the print of the called function, so it will not be straight away exceuted, rather the outputs will be stored in a stack space and they will be taken out in a LIFO (last in first out fashion)
// give hello charger {function call} 5 4 3 2 1 0 wordle okay away 
func mydef() int{
	for i := 0 ; i<6 ;i++ {
		defer fmt.Println(i);
	}
	// defer fmt.Println("cheese")
	return 0;
}
