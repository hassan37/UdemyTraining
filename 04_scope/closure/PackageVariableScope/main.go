package main

import "fmt"

var x int = 32

func main() {
	fmt.Println(increment())
}

func increment() int {
	x++;
	return x
}