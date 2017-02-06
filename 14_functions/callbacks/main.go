package main

import "fmt"

func main() {
	numbers := [] int{1, 2, 3}
	callbackExample(numbers, print)
	callbackExample(numbers, increment)
	fmt.Println()
	callbackExample(numbers, print)
	callbackExample(numbers, decrement)
	fmt.Println()
	callbackExample(numbers, print)
}

func callbackExample(numbers []int, callback func(x int) int) {
	for k, v := range numbers {
		numbers[k] = callback(v)
	}
}


func increment(x int) int {
	x++
	return x
}

func decrement(x int) int {
	x--
	return x
}

func print(x int) int {
	fmt.Print(x)
	return x
}