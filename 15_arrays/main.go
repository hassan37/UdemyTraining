package main

import "fmt"

/**
Go's arrays are values.
An array variable denotes the entire array;
it is not a pointer to the first array element (as would be the case in C).
This means that when you assign or pass around an array value you will make a copy of its contents
 */
func main() {
	//lengthIsPartOfArrayTypeDemo()
	//arraysArePassedAsValues()
	passAddressOfArrayIfModificationIsRequiredToBeDoneThroughFunction()
}


func lengthIsPartOfArrayTypeDemo() {
	var x [5]int
	var y [10]int

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	//fmt.Println(x == y) // ERROR: 15_arrays\main.go:15: invalid operation: x == y (mismatched types [5]int and [10]int)
	//x = y // ERROR: cannot use y (type [10]int) as type [5]int in assignment
}

func arraysArePassedAsValues() {
	var x [5]int
	fmt.Println("before:")
	fmt.Println("Address of x is: ", &x)
	fmt.Println("Value of x is: ", x)
	arrayPassedAsValue(x)
	fmt.Println("After:")
	fmt.Println("Value of x is: ", x)
}

func arrayPassedAsValue(x [5]int) {
	fmt.Println("Address of x is: ", &x)
	fmt.Println("Value of x is: ", x)

	for i := 0; i < len(x); i++ {
		x[i] = i + 100
	}
	fmt.Println("Updated Array x is: ", x)
}

func passAddressOfArrayIfModificationIsRequiredToBeDoneThroughFunction() {
	var x [5]int
	fmt.Println("before:")
	fmt.Println("Address of x is: ", &x)
	fmt.Println("Value of x is: ", x)
	addressOfArrayPassed(&x)
	fmt.Println("After:")
	fmt.Println("Value of x is: ", x)

}


func addressOfArrayPassed(x *[5]int) {
	fmt.Println("Address of x is: ", &x)
	fmt.Println("Value of x is: ", x)

	for i := 0; i < len(x); i++ {
		x[i] = i + 100
	}
	fmt.Println("Updated Array x is: ", x)
}
