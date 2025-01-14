package main

import "fmt"

func main() {
	ch := make(chan int)

	// IIFE funciton 
	go func() {
		ch <- 42 // This will block until the main goroutine is ready to receive
		fmt.Println("Sent 42")
	}()
	go func(){
		val :=<-ch
		fmt.Println("val recieved",val) 
	}()

	fmt.Println("Waiting to receive...")
	// value := <-ch // This will block until a value is available
	// fmt.Println("Received:", value)

}
