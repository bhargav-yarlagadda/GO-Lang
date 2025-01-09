package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5) // slice now contains [1, 2, 3, 4, 5]

	// creating of a slice
	slice = make([]int, 3, 5)
	fmt.Println(slice) 
	var newSlice[]int = make([]int,0,5)
	fmt.Println(newSlice)
	var newSlices []int;
	// or 
	// anotherSlice := make([]int)
	fmt.Println(newSlices)

}