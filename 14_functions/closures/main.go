package main

import "fmt"

//Closures Best Explained [https://dmitryfrank.com/articles/js_closures]
func main() {
	outerInnerScopeBoundaries()
	globalScope()
	limitingScopeByInnerFunctionExpressionDemo()
	limitingScopeByClosuresDemo()
}

//--------------------------- 1. OUTER INNER SCOPE

func outerInnerScopeBoundaries() {
	fmt.Println("***Outer Inner Scope Demo***")
	x := 42
	{
		fmt.Println("X is accessible in inner Scope: ", x);
		y := "inner scope variable";
		fmt.Println("cannot access y from outside of the scope: " + y);
	}
	//fmt.Println("cannot access y from outside of the scope: " + y);
}

//--------------------------- 2. Global Scope

var x int
func incrementByGlobalScope() int {
	x++
	return x
}

func globalScope() {
	fmt.Println("\n***Global Scope Demo***")
	fmt.Println(incrementByGlobalScope())
	fmt.Println(incrementByGlobalScope())
}

//--------------------------- 3. Limiting Scope Demo By Inner Function Expression

func limitingScopeByInnerFunctionExpressionDemo() {
	var x int
	increment := func() int {
		x++
		return x
	}
	fmt.Println("\n***Limiting Scope Demo By Inner Function Expression Demo***")
	fmt.Println(increment())
	fmt.Println(increment())

	//Now i cannot have multiple function instances to anonymous function defined above
	//every time i want to have a func expression of increment, i will rewrite anonymous function
}

//--------------------------- 4. Limiting Scope Demo By Closures

func limitingScopeByClosuresDemo() {
	fmt.Println("\n***Limiting Scope Demo By Closures Demo***")
	increment1 := incrementByClosures();
	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())

	increment2 := incrementByClosures();
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func incrementByClosures() func() int {
	var x int
	increment := func() int {
		x++
		return x
	}
	return increment
}
