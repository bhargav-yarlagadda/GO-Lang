package main
import "fmt"

// sturcts are non primitive data types to store key values pairs, similar to dicts in py or objetct models in js
// syntax
type Person struct{
	name string;
	age int
}


// structs can be nested  
type Human struct{
	evolution string;
	being Person
}


func main(){

	// declaring a struct
	guy := Person{}
	student := Person{name:"bhargav",age:18}
	
	fmt.Println("stuct is ",guy)
	fmt.Println("name is ",guy.name)
	fmt.Println("age is ",guy.age)


	// accessing a attribute
	student.name= "venkata Bhargav"
	
	fmt.Println(student)
}
