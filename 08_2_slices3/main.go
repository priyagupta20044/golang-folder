package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices lecture")

	marks := make([]int,3)
	marks[0]=1 
	marks[1]=3
	marks[2]=4
	fmt.Println(marks)
	var courses = []string {"p","q","r","s","t"}
	courses = append(courses,"u")

	// removing index 3 from the courses
	index := 3
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}

