package main

import (
	"bufio"
	"fmt"
	"os"
)

 func main() {
	fmt.Println("user input function go:")
     
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter your input here anything you thinking of :")
    
	// comma ok || error ok
    input,_:=reader.ReadString('\n')
	fmt.Println("your input printed by this program is :",input)
	

}