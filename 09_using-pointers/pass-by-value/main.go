package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println("x | Address ", x, " | ", &x)
}

func zeroPointer(x *int) {
	*x = 0
	fmt.Println("x | Address ", x, " | ", &x)
}

func main() {
	a := 5
	fmt.Println("a | Address ", a, " | ", &a)
	zero(a)
	fmt.Println("a | Address ", a, " | ", &a)
	zeroPointer(&a)
	fmt.Println("a | Address ", a, " | ", &a)

}
