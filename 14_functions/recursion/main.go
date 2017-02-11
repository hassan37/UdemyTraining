package main

import "fmt"

func main() {
	fmt.Println("X is:", 0, "| Next Fibonacci is:", nextFibonacci(0));
	fmt.Println("X is:", 1, "| Next Fibonacci is:", nextFibonacci(1));
	fmt.Println("X is:", 2, "| Next Fibonacci is:", nextFibonacci(2));
	fmt.Println("X is:", 3, "| Next Fibonacci is:", nextFibonacci(3));
	fmt.Println("X is:", 4, "| Next Fibonacci is:", nextFibonacci(4));
	fmt.Println("X is:", 5, "| Next Fibonacci is:", nextFibonacci(5));
	fmt.Println("X is:", 6, "| Next Fibonacci is:", nextFibonacci(6));
	fmt.Println("X is:", 7, "| Next Fibonacci is:", nextFibonacci(7));
	fmt.Println("X is:", 8, "| Next Fibonacci is:", nextFibonacci(8));
	fmt.Println("X is:", 9, "| Next Fibonacci is:", nextFibonacci(9));
	fmt.Println("X is:", 10, "| Next Fibonacci is:", nextFibonacci(10));

}

func nextFibonacci(x int) int {
	if 0 == x || 1 == x {
		return 1
	}

	return nextFibonacci(x-1) + nextFibonacci(x-2)
}
