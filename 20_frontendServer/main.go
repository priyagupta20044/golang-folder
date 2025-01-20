package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	// "net/http"
)

func main() {
	fmt.Println("This lecture demonstrates the handling url from the server made using Javascript and express")
	var Url string = getUrlFromUser()
	readContentFromUrl(Url)
}
func readContentFromUrl(Url string) {
	response, err := http.Get(Url); 
	checkNilErr(err); 

	fmt.Println("The status of the response is : ",response.StatusCode); 
	fmt.Println("The body of the response is : ",response.Body); 

	databytes, err := io.ReadAll(response.Body); 
	checkNilErr(err); 
	fmt.Println("The actual content is:",string(databytes)); 

	defer response.Body.Close();
}
func getUrlFromUser() string {
	Reader := bufio.NewReader(os.Stdin)
	MyUrl, err := Reader.ReadString('\n')
	checkNilErr(err)
	return strings.TrimSpace(MyUrl);
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
