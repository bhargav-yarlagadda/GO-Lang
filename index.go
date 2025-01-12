package main
import "fmt"
func main(){
    ageMap := make(map[string] int)
    fmt.Println(ageMap)
    // appending into map
    ageMap["bhargav"] = 21
    ageMap["maggie"] = 30
    // deleting from map
    delete(ageMap,"maggie")
    //checking for existence
    age,exists := ageMap["ramu"]
    if(exists){
            fmt.Println("ramu exists and age is ",age)
    }else{
        fmt.Println("ramu does not exits")
    }
    for key,value := range ageMap{
        fmt.Println(key,value)
    }

    // we can nest maps but this makes it difficult to understand why while working instead we can use structs as Keys

    // Define a struct
    type Person struct {
        FirstName string
        LastName  string
    }
    peopleAges := make(map[Person]int)
    peopleAges[Person{FirstName: "Alice", LastName: "Smith"}] = 30
    peopleAges[Person{FirstName: "Bob", LastName: "Johnson"}] = 25
    fmt.Println(peopleAges)
}