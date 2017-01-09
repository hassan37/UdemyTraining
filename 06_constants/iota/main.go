package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E
	F
)

const (
	P = iota + 1
	Q
	R = iota
	S
	T = iota

)

func main() {
	fmt.Println(A); fmt.Println(B); fmt.Println(C);
	fmt.Println(D); fmt.Println(E); fmt.Println(F);
	fmt.Println()
	fmt.Println(P); fmt.Println(Q); fmt.Println(R); fmt.Println(S); fmt.Println(T);
}
