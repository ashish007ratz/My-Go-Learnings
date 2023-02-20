package main

import "fmt"

func main() {

	//pointers
	fmt.Println("learing pointers in go ")
	// declare a pointer
	var ptr *int
	fmt.Println("default value for ptr is:", ptr)

	myNum := 23
	var myPointer = &myNum
	fmt.Println("memory add of pointer is now", myPointer)
	fmt.Println("value of pointer is now", *myPointer)

	*myPointer = *myPointer * 2
	fmt.Println("value of pointer is now", myNum)
}
