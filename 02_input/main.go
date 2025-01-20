package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	number1, _ := reader.ReadString('\n')
	number2, _ := reader.ReadBytes('\n')

	fmt.Println("The numbers are : ", number1)
	fmt.Println("The numbers are : ", number2)
}
