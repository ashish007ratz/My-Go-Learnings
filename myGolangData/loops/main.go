package main

import "fmt"

func main() {
	fmt.Println("Loops in go lang")

	days := []string{"Sunday","Monday","Tuesday", "Wednesday","Thursday","Friday","Saturday"}

	fmt.Println(days)

	for d:=0;d<len(days); d++{
		fmt.Println(days[d])
	}

	// diff loops
	for i:= range days{
		fmt.Println(days[i])
	}

	// for each loop
	for index,day:= range days{
 fmt.Printf("index is %v and value is %v\n",index,day)
	}

//
rougeValue :=1
for rougeValue <10{
	if(rougeValue ==2){
		goto lco
	}
	if(rougeValue == 5){
		rougeValue ++
		continue
	}
	
	fmt.Println("Vlaue is: ", rougeValue)
	rougeValue ++
}

// go to statement
lco:
fmt.Println("learing go online.in")

}