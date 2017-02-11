package main

import "fmt"

func main() {
	open()
	defer close() // it does make sense to write close immediately after open. and
	process()
	save()
}

func open() {
	fmt.Println("opened")
}

func close() {
	fmt.Println("closed")
}

func process() {
	fmt.Println("processed")
}

func save() {
	fmt.Println("saved")
}