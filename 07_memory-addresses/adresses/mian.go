package main

import "fmt"

func main() {
	a := 2
	fmt.Println("a - ", a, "\nAdress of a - ", &a)
	fmt.Printf("Decimal Adress of a - %d", &a)
}
