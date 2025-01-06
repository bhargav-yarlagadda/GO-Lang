package main
import "fmt"
// syntax

func main(){
	var i int 
	i = 10
	fmt.Println("the value of i is ",i)
	// short hand
	variable:=10 
	fmt.Println("value of vaiable is",variable)
	// cannot change its type
	variable=23
	// cannot use format specifier in println
	floatNum :=2.6
	intNum :=int(floatNum)
	fmt.Printf("value of intNUm is %d",intNum)


	// constants
	const pi float32 = 2.14
	// we cant use short notation for consts
}