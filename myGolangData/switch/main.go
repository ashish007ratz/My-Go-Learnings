package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("switch case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1
	fmt.Println("value of dice is :",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open or take one step farward")
	         
	case 2:
		fmt.Println("turn 2 step a head")
         fallthrough
	case 3:
		fmt.Println("turn 3 step a head")

	case 4:
		fmt.Println("turn 4 step a head")

	case 5:
		fmt.Println("turn 5 step a head")

	case 6:
		fmt.Println("you can open and play one more chance")
     default :
	 fmt.Println("Unlucky you")

	}
}