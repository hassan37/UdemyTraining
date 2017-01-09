package main

import "fmt"

func main() {
	i:=0
	for {
		i++
		if i % 2 == 1 && i % 3 == 0 {
			fmt.Println(i)
		}
		if i > 50 {
			break
		}
	}
}
