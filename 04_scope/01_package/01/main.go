package main

import "fmt"

//scope of x is: Package Level
// Whole package can access x variable

var x int = 2

func main() {
	fmt.Println("Pkg Level Scope: x in main : ", y)
	foo()
}

func foo() {
	fmt.Println("Pkg Level Scope: x in foo : ", x)
}
