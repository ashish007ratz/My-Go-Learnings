package main

import "fmt"

func main() {
	fmt.Println("welcome to fucntion mian")
	newFuction()

	// add two number fucntion
	result := addTwo(20, 40)
	fmt.Println(result)
	// multi add
	fmt.Println(addMulti(1,23,4,5,56,6))

}

/// create new function
func newFuction() {
	fmt.Println("Now you have entered new function")
}

// fucntion with arguments
func addTwo(valueOne int, valTwo int) int {
	println("vale one is", valueOne)
	println("vale two is", valTwo)
	return (valueOne + valTwo)
}

func addMulti(values ...int)int{
	sum := 0
    for i:= range values{
		sum =sum+values[i]
	}   
	return sum
}
