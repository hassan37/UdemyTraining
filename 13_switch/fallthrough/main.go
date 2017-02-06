package main

import "fmt"

func main() {
	x := "hello"
	fmt.Print("Enter: ")
	fmt.Scan(&x)
	switch x {
	case "hello":
		fmt.Println("Hello")
		fallthrough
	case "go":
		fmt.Println("GoLang")
	case "world":
		fmt.Println("World")
	default:
		fmt.Println("defworld")
	}

}
