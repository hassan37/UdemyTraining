package main

import "fmt"

func main() {
	s := greetings("Hassan", "Welcome to GoLang")
	fmt.Println(s)
	x,y := multipleReturns("Hassan: ", "Welcome to GoLang")
	fmt.Println(x,y)
}

func greetings(name, message string) string {
	return fmt.Sprint(name, message)
}

func multipleReturns(name, message string) (string, string) {
	return fmt.Sprint(name, message), fmt.Sprint(message, name)
}
