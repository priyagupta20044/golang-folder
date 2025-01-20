package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2025, time.January, 24, 15, 44, 49, 11, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
	my_birthday := time.Date(2004,time.July,31,2,25,34,11,time.UTC); 
	fmt.Println(my_birthday.Format("01-02-2006 15:04:05 Monday"));

}
