package main

import "fmt"

func main() {

	//Basic for loop with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//A classic inital/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//For without a condition loops repeadly until
	//broken by a break or return for the encolsing function
	for {
		fmt.Println("loop")
		break
	}
}