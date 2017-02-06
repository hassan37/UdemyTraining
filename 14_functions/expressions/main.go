package main

import "fmt"

func main() {
	funcExprGreetings()
	returnFuncType()
}


//1. function expression: assigning a function to a variable:
// [ONLY anonymous func can be used to do that. the type will be 'func()' type
//functions as variable
func funcExprGreetings() {
	greetings := func(name, message string) {
		fmt.Println(name, message)
	}
	greetings("Hassan", "Welcome to GoLang")
}

//2. functions as return types
// closures
func returnFuncType() {
	greeter := returnFuncTypeGreeter()
	fmt.Printf("greeter type: %T\n", greeter)
	fmt.Println(greeter())
}

func returnFuncTypeGreeter() func() string {
	return func() string {
		return "Hello Hassan! Welcome to Clousres."
	}
}
