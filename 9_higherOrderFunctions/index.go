package main

import "fmt"

func main() {
	// we are now going to see higher order functions
	// what are higher order functions ???
	// functions that either return functions or take another functions as an argument
	// these are also known as a first class functions
	// lest understand with an example
	fmt.Println("examples of lower order functtions")
	fmt.Println(add(2,3),"is a lower order function.")
	fmt.Println(mult(2,3),"is a lower order function.")

	fmt.Println("example fo higher order functions that accepts functions as argument")
	fmt.Println(aggregate(3,5,1,add),"is a higher order function")
	fmt.Println(aggregate(10,2,3,mult),"is a higher order function")

	fmt.Println("example fo higher order functions that returns  functions as a value")
	fn:=higherOrderFuntion(1,2)
	fmt.Println(fn(3,4))
	// op should be 1 + 2 + 3 +4

}
func add(x int, y int) int {
	return x + y
}
func mult(x int, y int) int {
	return x * y
}
func aggregate(a int, b int, c int, fn func(int, int) int) int {
	return fn(fn(a, b), c)
}
func higherOrderFuntion(a int,b int) func(int,int) int{
	adder := func(x, y int) int {
		return a + b + x + y
	}
	return adder

}