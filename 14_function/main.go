package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This lecture is on functions")
	greetings()

	fmt.Println("Enter the values to be added:")

	reader := bufio.NewReader(os.Stdin)
	int1, _ := reader.ReadString('\n')
	int2, _ := reader.ReadString('\n')
	int3, _ := strconv.ParseInt(strings.TrimSpace(int1), 32, 32)
	int4, _ := strconv.ParseInt(strings.TrimSpace(int2), 32, 32)

	ans := adder(int3,int4);

	fmt.Printf("The sun of %d and %d is: %d\n", int3, int4, ans)

	// fmt.Println("Enter number of values to be addded: "); 
	prores := proAdder(1,2,4,5,3);
	fmt.Printf("The total sum is : %d" ,prores); 


}
func greetings() { // functions cannot be declared inside another functio but outside
	fmt.Println("Namaste")
}
func adder(a int64, b int64) int64 {
	ans := a + b
	return ans
}
func proAdder(values ... int) int {
	total := 0 ; 
	for _,iter := range values {
		total += iter;
	}

	return total;
}

