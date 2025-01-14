package main

import (
	"fmt"
	"time"
)

func greet(i int) {
	fmt.Println("hello from routine",i)
}
func main() {
	go greet(0)
	fmt.Println("hello from main")
	time.Sleep(time.Second)
	for i := 1; i <= 10; i++ {
		go greet(i)
		time.Sleep(time.Second)
	}
	

}

// in the below snippets we are not waiting so the routines run in background , before completion of all routines main prog executes parallely
// thus all then statements will next be printed , only one or two will be printed

// func main() {
// 	// Start 10 goroutines without waiting
// 	for i := 1; i <= 10; i++ {
// 		go greet(i)
// 	}

// 	// The main function can finish without waiting for the goroutines.
// 	fmt.Println("Main function exits")
// }
