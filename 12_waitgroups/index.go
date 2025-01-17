package main

import (
	"fmt"
	"net/http"
	"sync"
)


var wg  sync.WaitGroup //pointer
func main(){

	websiteList :=[]string{
		"https://youtube.com",
		"https://go.dev",
		"https://leetcode.com",
		"https://fb.cim",



	}

	// we are synchrnously running them they take a while to run
	// for _,web := range websiteList{
	// 	getStatusCode(web)
	// }
	
	for _,web := range websiteList{
		go	getStatusCode(web)
		wg.Add(1)

	}
	// this is responsible for not exiting the main method
	wg.Wait()


}

func getStatusCode(endpoint string) {
	defer wg.Done() // this method will be excecuted at the end of getStatus code 
	// even though its called at the top it will be only executed at the end of this function call. 
	res,err:=	http.Get(endpoint)	
	if err!=nil{
		fmt.Println("error in endpoint, please ensure the endpoint is valid")
	}else{
		fmt.Printf("%v status code for website %v \n",res.StatusCode,endpoint)
	}
}
