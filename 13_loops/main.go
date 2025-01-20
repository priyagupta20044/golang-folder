package main

import "fmt"

func main() {
	fmt.Println("This is the lecture on loops")
	days := []string {"Sunday","tuesday","Friday","Sunday"}; 

	for d := 0;d< len(days) ;d++{
		fmt.Println(days[d]); 
	}

	for i:= range days{
		fmt.Println(days[i]);
	}

	for idx, day := range days{
		fmt.Printf("The index is %d and day is :%s\n",idx,day);
	}

	init := 12; 

	for init <20{
		if init == 17 {
			continue;
		}
		if (init==15){
			goto GOTO;
		}
		fmt.Println(init); 
		init ++;
	}

	GOTO : fmt.Println("You jumped on the goto statement")
}
