package main

import "fmt"

func main() {
	fmt.Println("*** Integer Pass By Value Demo ***")
	demoPassByValueInt()

	fmt.Println()

	fmt.Println("*** String Pass By Value Demo ***")
	demoPassByValueString()
}

func demoPassByValueInt() {
	incrementDemoByInt := func() {
		increment := func(a int) {
			a++
			fmt.Println("Received: ", a, "Address:", &a, "Incremented Value:", a)
		}
		x := 5
		fmt.Println("Before: ", x, "Address:", &x)
		increment(x)
		fmt.Println("After: ", x)
	}
	incrementDemoByInt()

	incrementDemoByIntPointer := func() {
		increment := func(a *int) {
			*a++
			fmt.Println("ReceivedPtr:", a, "AddressOfPtr:", &a, "Incremented Value:", *a)
			y := 10
			a = &y
			fmt.Println("NewPtr:", a, "AddressOfPtr:", &a, "Incremented Value:", *a)
		}
		x := 5
		xPtr := &x
		fmt.Println("BeforePtr: ", xPtr, "Address:", &xPtr, "ValueAtPtr:", *xPtr)
		increment(xPtr)
		fmt.Println("AfterPtr: ", xPtr, "Address:", &xPtr, "ValueAtPtr:", *xPtr)
	}
	incrementDemoByIntPointer()


}


func demoPassByValueString() {
	incrementDemoByString := func() {
		increment := func(a string) {
			a = "a"
			fmt.Println("Received: ", a, "Address:", &a, "Incremented Value:", a)
		}
		x := "x"
		fmt.Println("Before: ", x, "Address:", &x)
		increment(x)
		fmt.Println("After: ", x)
	}
	incrementDemoByString()


	incrementDemoByStrPointer := func() {
		increment := func(a *string) {
			*a = "a"
			fmt.Println("ReceivedPtr:", a, "AddressOfPtr:", &a, "Incremented Value:", *a)
			y := "y"
			a = &y
			fmt.Println("NewPtr:", a, "AddressOfPtr:", &a, "Incremented Value:", *a)
		}
		x := "x"
		xPtr := &x
		fmt.Println("BeforePtr: ", xPtr, "Address:", &xPtr, "ValueAtPtr:", *xPtr)
		increment(xPtr)
		fmt.Println("AfterPtr: ", xPtr, "Address:", &xPtr, "ValueAtPtr:", *xPtr)
	}
	incrementDemoByStrPointer()


}
