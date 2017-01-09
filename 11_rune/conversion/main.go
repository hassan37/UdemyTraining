package main

import "fmt"

func main() {
	a := []byte("Hello")
	fmt.Println(a)
	bRune := []byte("世")
	fmt.Println(bRune)
	i := 0
	for {
		fmt.Printf("%v | %v | %v\n", i, string(i), []byte(string(i)))
		i++
		if i > 1000 {
			break
		}
	}
}
