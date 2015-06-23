package main

import "fmt"
import "time"

func main() {

	//Her is a basic switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}

	//You can use commas to seperate multiple expressions
	//in the same case statement. Also note the optional
	//default case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	//switch wothout an expression is an alternative way to
	//express if/else logic. Here we also show how the case expressions
	//can be non-constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}