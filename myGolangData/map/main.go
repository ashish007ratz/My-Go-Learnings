package main

import "fmt"

func main() {
	fmt.Println("learing maps in goloang")

	dataMap := make(map[string]string)
	dataMap["data1"] = "class1"
	dataMap["data2"] = "class2"
	dataMap["data3"] = "class3"
	dataMap["data4"] = "class4"
	dataMap["data5"] = "class5"

	fmt.Println("values of map is : ", dataMap)
	fmt.Println("data1 of map is : ", dataMap["data1"])

	// remove data from map
	delete(dataMap,"data1")
	fmt.Println("values of map is : ", dataMap)

	// loops in map
	for key,value := range dataMap{
		fmt.Printf("for key %v, value is %v\n",key,value)
	}
}
