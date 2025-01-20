package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in go language")

	var arr[7] string; 
	arr[0]="Apple"
	arr[3]="Plums"
	arr[4]="Cherry"
	arr[6]="Banana"
	fmt.Println("The array is: ",arr)

	fmt.Println("The length of the array: ",len(arr))
	
	var myArr = [5] string {"potato","onion","carrot","brinjal","bottle guard"}
	fmt.Println("The vegetable array is:",myArr)
}
