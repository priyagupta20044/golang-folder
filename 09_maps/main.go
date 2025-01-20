package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps lecture")

	my_language_map := make(map[string]string)
	my_language_map["js"] = "javascript"
	my_language_map["cpp"] = "cplusplus"
	my_language_map["html"] = "hypertextmarkuplanguage"

	my_class_map := make(map[int]string)
	my_class_map[0] = "priya"
	my_class_map[1] = "maths"
	my_class_map[2] = "coding"
	my_class_map[3] = "development"

	fmt.Println("The list of all languages",my_language_map)

	delete(my_class_map,1)
	fmt.Println(my_class_map)
}
