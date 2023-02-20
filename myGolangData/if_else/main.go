package main

import "fmt"

func main(){

	fmt.Println("Hello if else in go")

	count :=1

	if (count == 1){
		fmt.Println("count is ok not out of limit")
	
	}else {
		fmt.Println("count is not ok printing else statement")
	}

	/// special syntax
	if num :=3; num ==4{
		fmt.Println("if condition")
	}else{fmt.Println("else condition")}
}