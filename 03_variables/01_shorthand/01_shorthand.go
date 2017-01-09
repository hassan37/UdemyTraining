package main

import "fmt"

func main() {
	//Go has a shorthand for variable declarations using := instead of =:
	//x := "Hello World"
	//This is equivalent to:
	//var x = "Hello World"
	//This shorthand is only allowed to be used within functions.

	a := 1
	b := "hello go lang"
	c := 20.5
	d := true
	e := 'i'
	f := `Do you like my hat?`

	//%v: prints values in their default formats
	// for example: for bool type, it inserts %t
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

	fmt.Println()

	//%T: prints the type of value
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}

