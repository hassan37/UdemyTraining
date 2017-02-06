package main

import "fmt"

func main() {
	//comma separated arguements are being passed to varadic parameters
	sum := getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum)

	//defining slice
	nums := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("\nType of Slice nums: %T \n", nums)

	//slice is being passed as VARIADIC ARGUMENT
	avg := getAvg(nums...)
	fmt.Println(avg)
}

func getSum(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}

	return sum
}

func getAvg(nums ...float64) float64 {
	fmt.Printf("\nType of Varidic Param Nums: %T \n", nums)

	//slice is being passed as VARIADIC ARGUMENT
	sum := getSum(nums...)

	return sum / float64(len(nums))
}