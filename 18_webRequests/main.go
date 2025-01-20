package main

import (
	"fmt"
	"io"
	"net/http"
)

var url = "https://www.geeksforgeeks.org/java/"

func main() {
	fmt.Println("Welcome to the lecture on web server request handling")
	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("The type of the response is: %T", response)
	fmt.Println(response)

	defer response.Body.Close() // callers responsibility to close the body of the response

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("The content of the website is:", string(databytes))
}
func checkNilErr(Err error) {
	if Err != nil {
		panic(Err)
	}
}
