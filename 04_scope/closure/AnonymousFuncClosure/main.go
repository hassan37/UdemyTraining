package main

import "fmt"

func main() {
	x := 32
	//FUnction Expression: assigning a function to a variable
	//Anonymous function: limits/closures the access level to a variable
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
