package main

import "fmt"

func main() {
	intSwitch()
	floatSwitch()
	boolSwitch()
	stringSwitch()
}

func stringSwitch() {
	a := "hello"
	switch a {
	case "hello":
		fmt.Println("a is hello")
	}
}

func boolSwitch() {
	a := true
	switch a {
	case true:
		fmt.Println("a is true")
	}
}

func floatSwitch() {
	a := 1.5
	switch a {
	case 1.5:
		fmt.Println("a is 1.5")
	default:
		fmt.Println("a is not 1.5")
	}
}

func intSwitch() {
	a := 1
	switch a {
	case 1:
		fmt.Println("a is 1")
	default:
		fmt.Println("a is not 1")
	}
}
