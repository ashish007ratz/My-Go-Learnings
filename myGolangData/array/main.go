package main

import "fmt"
func main (){
	fmt.Println("learing array on go")

	var listName [4]string

	listName[0]="Ashish"
	listName[1]="Amit"
	listName[2]="sumit"
	
	fmt.Println("List of name is: ", listName)
	fmt.Println("list length is :",len(listName))
    
	var dataList =[3]string{"data1","data2","data3"}
    fmt.Println("veg list is :",dataList)
}