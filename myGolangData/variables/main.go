package main

import "fmt"

func main() {
	var username string = "ashish"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	//bool
	var isBool bool = false
	fmt.Println(isBool)
	fmt.Printf("variable type is : %T",isBool)

	//int
	var smallInt uint8 =255//uint8(0-255)uInt16(0-65535),uint32(0-4294967295)uint64(18446744073709551615)
	fmt.Println(smallInt)
	fmt.Printf("variable type is : %T",smallInt)

	//float
	var floatNum float32 =1002.432
	fmt.Println(floatNum)
	fmt.Printf("variable type is : %T",floatNum)

     // default values and some aliases
	 var number int
	 fmt.Println(number)
	fmt.Printf("variable type is : %T",number)

	// imlicit type
	var data ="my name is ashish"
	fmt.Println(data)
	
	//no var style
	noVar := 1000
	fmt.Println(noVar)

	// public variable
	fmt.Println(Token)
	fmt.Printf("variable type is : %T",Token)

}
const Token string ="hello my name is ashsih"