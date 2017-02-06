package main

import "fmt"

type Contact struct {

	greeting string
	name string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("it's an int")
	case string:
		fmt.Println("it's an String")
	case Contact:
		fmt.Println("it's Contact Type")
	}
}

func main() {
	switchOnType(7)
	switchOnType("Hassan")
	t := Contact{"Hello", "Hassan"}
	switchOnType(t)
}
