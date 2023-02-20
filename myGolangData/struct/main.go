package main

import "fmt"

func main(){

	fmt.Println("this is my struct program")

	//no inheritance in golang ; no super ,parent
    ankit:= Student{"Ankit",12,"class1","Dehradun"}
    fmt.Println(ankit)
	fmt.Printf("Ankit details are: %v\n",ankit)  
}

type Student struct{
	Name string
	Age int
	Class string
	City string 
}