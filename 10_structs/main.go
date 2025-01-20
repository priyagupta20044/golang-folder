package main

import "fmt"

func main() {
	fmt.Println("Strcutures in Golang")
	// there is no inheritence; no super or parent class
	Priya := User{"Priya","priya.gupta20044@gmail.com",true,21};
	fmt.Println("User name is:",Priya.Name);
	fmt.Println("Age of priya is:",Priya.Age);
	fmt.Println("Email of Priya is:",Priya.Email); 
	fmt.Println("Priya is logged in:",Priya.Status);
	fmt.Printf("The details of user are:%+v\n",Priya)

	fmt.Printf("name is :%v and email is :%v",Priya.Name,Priya.Email);
}

type User struct { // classes with capitalization can be used to export into other classes 
	// all the members with the small initials will not be exported whereas all the elements with capital initals can be exported
	Name   string
	Email  string
	Status bool
	Age    int
}
