package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learing slices")

	var dataList = []string{"data1", "data2"}
	fmt.Printf("Type of dataList is %T\n", dataList)

	/// append slices
	dataList = append(dataList, "data3", "Data4")
	fmt.Println(dataList)

	// seprate part of your slice
	dataList = append(dataList[1:])
	fmt.Println(dataList)

	scors := make([]int, 4)
	scors[0] = 10
	scors[1] = 20
	scors[2] = 30
	scors[3] = 40

	scors = append(scors, 50, 60, 70, 80)
	fmt.Println(scors)
	sort.Ints(scors)
	fmt.Println(scors)

	// remove value from slice based on index
	var classes =[]string{"class1","class2","class3","class4","class5","class6","class7"}
    fmt.Println(classes)
	var rmv int=2
	classes =append(classes[:rmv],classes[rmv+1:]...) 
    fmt.Println(classes)
	
}
