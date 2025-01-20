package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=rgy54edth"

func main() {
	fmt.Println("Welcome to the lecture on URLs")
	fmt.Println(myUrl)

	// parsing
	parsedUrl, err := url.Parse(myUrl)
	checkNilErr(err);
	fmt.Println("The Parsed ULR is:",parsedUrl)

	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.Port())
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.RawQuery)

	qparams := parsedUrl.Query(); 
	fmt.Printf("The type of the q params is :%T\n",qparams);

	// fmt.Println(qparams["coursename"])

	for t,val := range qparams{
		fmt.Printf("The value of %s the param is :%s\n",t,val);
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "learncss", 
		RawPath : "user=Priya",
	}

	anotherUrl := partsOfUrl.String(); 

	fmt.Println("The new url is :",anotherUrl)
}
func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}
