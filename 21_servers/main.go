package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("This is handling web servers in Go")
	handleGetRequest()
	handlePostRequest()
	handlePostFormRequest()
}
func handleGetRequest() {
	const Url = "http://localhost:3001"

	// getting the response
	response, err := http.Get(Url)
	checkNilErr(err)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println("The status code of the response is:", response.StatusCode)
	fmt.Println("The length of the response is", response.ContentLength)

	var stringConverter strings.Builder
	stringConverter.Write(content)
	fmt.Println("The content inside the response body is :", stringConverter.String())

}
func handlePostRequest() {
	const Url = "http://localhost:3001/post"
	// First, generating the request that we will feed into the post
	requestBody := strings.NewReader(`
		{
			"flat":"141",
			"tower":"Ruby",
			"society":"Gardenia Glamour"
		}
	`)

	// Second, add this request to the post by using http
	response, err := http.Post(Url, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()

	// Return the response recieved into content
	content, _ := io.ReadAll(response.Body)

	// print the string version of the response. The content is still in thr bytes format so we convert it to string
	var stringConverter strings.Builder
	stringConverter.Write(content)
	fmt.Println("The returned response is:", stringConverter.String())
}
func handlePostFormRequest() {
	const MyUrl = "http://localhost:3001/postform"
	fmt.Println("Welcome to the postFrom Request End point")

	// generate raw data (form data)
	RequestFormData := url.Values{};
	RequestFormData.Add("Subject","DBMS"); 
	RequestFormData.Add("Dept","CSE"); 
	RequestFormData.Add("Prof","DKT"); 
	RequestFormData.Add("Marks","97");

	// getting response from the end point 
	reponse,err := http.PostForm(MyUrl,RequestFormData);
	checkNilErr(err); 

	// Conversion into readable form 
	content,_ := io.ReadAll(reponse.Body); 
	var stringConverter strings.Builder; 
	stringConverter.Write(content);
	fmt.Println("The content body is:",stringConverter.String())

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
 