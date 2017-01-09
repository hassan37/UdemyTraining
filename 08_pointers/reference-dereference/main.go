package main

import "fmt"

func main() {
	a := "hello"
	b := a
	fmt.Println("a - ", a, "\nAdress of a - ", &a)
	fmt.Println("b - ", b, "\nAdress of b - ", &b)
	pa := &a
	fmt.Println("pa - ", pa, "\nValue at pa - ", *pa)
}
