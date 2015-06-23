package main

import "fmt"

func main() {

	//var declares one or more variables
	var a string = "initial"
	fmt.Println(a)

	//You can declare more than one variable at a time
	var b, c int = 1, 2
	fmt.Println(b, c)

	//Go will infer the type of the infered variable
	var d = true
	fmt.Println(d)

	//Variables declared without a coresponding initilization
	//are zero-valued
	var e int
	fmt.Println(e)

	//The := syntax is shaorthand for declaring and initilizing a variable
	f := "short"
	fmt.Println(f)
}