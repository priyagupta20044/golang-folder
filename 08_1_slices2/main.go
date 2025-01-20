package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the slices lecture")
	var slice1 = []string{"h", "i"}
	fmt.Printf("The type of the slice is %T\n ", slice1)

	slice1 = append(slice1, "j", "k", "l")
	fmt.Println("The new value of slice 1:", slice1)

	slice1 = append(slice1[1:3])
	fmt.Println(slice1)

	highScore := make([]int, 5)
	highScore[3] = 3

	highScore = append(highScore, 3, 4, 5, 6)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

}
