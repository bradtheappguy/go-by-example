package main

import "fmt"
import "math"

//Const declares a constnat value
const s string = "constant"

func main() {
	fmt.Println(s)

	//A const statement can appear anywhere a var can
	const n = 500000000

	//Constant expressions perform arthemetric with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	//A numeric constant has no type until it's given one by an exclusive cast
	fmt.Println(int64(d))

	//A number can be given a type by using it in a context that
	//requires one, such as a variable assignment of function call.
	fmt.Println(math.Sin(n))
}