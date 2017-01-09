package main

import "fmt"

func main() {
	x := 32
	{
		fmt.Println(x)
		y := "Declared Variable inside block"
		fmt.Println(y)
	}
	//fmt.Println(y) //not accessible
}
