package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to our new Pizza app")
	fmt.Println("Enter the rating of our app")

	newAppRating, _ := reader.ReadString('\n')
	// converting to int

	convertedRating, err := strconv.ParseFloat(strings.TrimSpace(newAppRating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("We increaed the rating to: ", convertedRating + 1)
	}
}
