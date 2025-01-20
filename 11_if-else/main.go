package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("GoLang Program for If-Else")
	reader := bufio.NewReader(os.Stdin)

	loginCount, err1 := reader.ReadString('\n')
	if err1 != nil {
		fmt.Println("err1")
	}
	converted_loginCount,_ := strconv.ParseFloat(strings.TrimSpace(loginCount),64); 

	var result string
	if converted_loginCount < 10 {
		result = "normal User"
	} else{
		result ="suspicious"
	}
	fmt.Println(result)
}
