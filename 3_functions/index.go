package main

import "fmt"

// function syntax 
// func fnName(arg1 datatype,arg2 datatype) datatypeOfFn {

// 	return 
// }

func add(x int,y int) int{

	return (x+y)
}


// call back ,here add is the name of callback func(int,int) is the type of function int is the return type
func addTwo(add func(int,int) int ,num int) int {
	return add(num,num) + num
}



func main(){
	fmt.Println("staring with functions")
	fmt.Println("a simple function ",add(1,2))
	fmt.Println("called the function that takes a function as a input",addTwo(add,10))
}

