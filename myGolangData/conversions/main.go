package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to our rating app")
	fmt.Println("rate your self out of 10")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("your input is:",input )

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(numRating+1)
	}
}