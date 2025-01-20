package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the lecture on file handling")

	content := "This is the content of the file which needs to go into it"

	file, err := os.Create("./my_dir_file_important.txt")
	// if err != nil {
	// 	panic(err) // panic is a keyword which is used to stop the execution of the program and print the error
	// }
	checkNilErr(err); 

	length, err := io.WriteString(file, content) // this function is used to put the content (from content) to the specified location (file here) and also returns the length of the file, ie the number of characters in it
	checkNilErr(err);
	fmt.Println("The length of the file is :", length)
	readFile("./my_dir_file_important.txt")
	defer file.Close()
}

func readFile(filename string) {
	bytestring , err := ioutil.ReadFile(filename);
	checkNilErr(err);
	fmt.Println("The data inside the file is :\n",string(bytestring));
}
func checkNilErr(err error){
	if err != nil {
		panic(err) // panic is a keyword which is used to stop the execution of the program and print the error
	}
}