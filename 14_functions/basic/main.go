package main

import "fmt"

func main() {
	greetings("Hassan", "Welcome to GoLang")
}

// func [receivers] name([parameters]) return;


func greetings(name, message string) {
	fmt.Println(name, message)
}
