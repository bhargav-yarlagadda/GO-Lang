package main

import "fmt"


// we cannot just pass the array we have to pass the exact size too
func printArr(arr [3]int){
	fmt.Println(arr,len(arr))
}

// returning array
func returnArr()[3]int{
	return [3]int{1,2,4}
}
func main() {
	// arrays --> contigious block,same data type ,fixed len
	// declaRATION
	var arr [5]int
	arr[1] = 10
	fmt.Println(arr)
	arr2 := [5]int{1,3,4}
	fmt.Println(arr2)
	newArr := [3]int{1,2,3}
	printArr(newArr)
	fmt.Println("return type")
	printArr(returnArr())

	fmt.Println("we can use normal indexes to access and modify arrays")

}