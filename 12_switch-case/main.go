package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	dice_number := rand.Intn(6) + 1

	switch dice_number {
	case 1:
		fmt.Println("Dice number is 1 and you can open")
	case 6:
		fmt.Println("Dice number is 6 and you can roll the dice again");
	default: 
	fmt.Printf("Dice number is %d\n",dice_number);
	}

}
