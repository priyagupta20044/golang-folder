package main

import "fmt"

func main() {
	fmt.Println("This is a lecture on Methods")
	class1 := My_class{"Priya",20,true,"priya.gupta20044@gmail.com"}; 
	class1.GetStatus();
	fmt.Println("the new mail is: ",class1.ChangeMail());
	fmt.Printf("%+v\n",class1);
}

type My_class struct {
	Name   string
	Age    int
	Status bool
	Email string
	// oneAge int // non exportable
}
func (class My_class) GetStatus(){
	fmt.Println("The user is active:",class.Status);
}
func (class *My_class) ChangeMail() string{
	class.Email = "hello.dev.in"
	return class.Email
}